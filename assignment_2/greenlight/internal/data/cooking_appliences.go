package data

import (
	"time"
)

type CookingApplience struct {
	ID int64 `json:"id"`
	CreatedAt time.Time  `json:"-"`
	Name string `json:"name"`
	Year int32 `json:"year,omitempty"` 
	Material string `json:"material,omitempty"`
	Categories []string `json:"categories,omitempty"`
	Version int32 `json:"version"`
}
