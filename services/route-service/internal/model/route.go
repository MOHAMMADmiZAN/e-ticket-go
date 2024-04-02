// internal/model/route.go

package model

import (
	"time"
)

// Route represents the route data structure.
type Route struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
}
