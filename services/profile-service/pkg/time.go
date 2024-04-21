package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// ConvertUTCtoLocal formats a UTC time to "M/D/YYYY, hh:mm:ss AM/PM" format using the specified or local timezone.
func ConvertUTCtoLocal(utcTime time.Time) (string, error) {
	// Try to get timezone from the environment variable TZ, default to Local if not set
	timezone, err := GetTimezone()
	if err != nil {
		timezone = os.Getenv("TZ")
		if timezone == "" {
			timezone = "Local"
		}
	}

	// Load the specified timezone
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return "", fmt.Errorf("error loading time zone '%s': %v", timezone, err)
	}

	// Convert the UTC time to the specified local time
	localTime := utcTime.In(location)
	// Format the time in the desired layout
	layout := "1-2-2006, 3:04:05 PM MST"
	return localTime.Format(layout), nil
}

// ConvertTime tries to convert a given time to local time using a specified or local timezone and formats it; logs errors if any.
func ConvertTime(t time.Time) string {
	localTime, err := ConvertUTCtoLocal(t)
	if err != nil {
		fmt.Printf("Error converting time (%v) to local: %v\n", t, err)
		// Fallback to a simpler UTC format if there's an error
		return t.Format("2006-01-02 15:04:05 MST")
	}
	return localTime
}

type IPInfoResponse struct {
	Timezone string `json:"timezone"`
}

// GetTimezone fetches the current server's timezone based on its public IP.
func GetTimezone() (string, error) {
	token := os.Getenv("IPINFO_TOKEN")
	if token == "" {
		return "", fmt.Errorf("IPINFO_TOKEN is not set in the environment variables")
	}
	ipInfoURL := "https://ipinfo.io/json?token=" + token + "&fields=timezone"

	resp, err := http.Get(ipInfoURL)
	if err != nil {
		return "", fmt.Errorf("could not retrieve IP information: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("error closing response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-ok HTTP status from ipinfo.io: %s", resp.Status)
	}

	var ipInfo IPInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&ipInfo); err != nil {
		return "", fmt.Errorf("could not decode IP information JSON: %v", err)
	}

	if ipInfo.Timezone == "" {
		return "", fmt.Errorf("timezone information is missing in the response")
	}

	return ipInfo.Timezone, nil
}
