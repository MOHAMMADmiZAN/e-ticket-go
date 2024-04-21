package pkg

import "encoding/base64"
import "crypto/rand"

func GenerateToken() (string, error) {
	// Define the size of the token in bytes (32 bytes equals 256 bits)
	b := make([]byte, 32)

	// Read random bytes; if there's an error, return it
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	// Return a URL-safe base64-encoded version of the token
	return base64.URLEncoding.EncodeToString(b), nil
}
