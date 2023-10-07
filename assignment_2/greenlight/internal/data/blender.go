package data

import (
	"time"

	"greenlight.alexedwards.net/internal/validator" // New import
)

type Blender struct {
	ID int64 `json:"id"`
	CreatedAt time.Time  `json:"-"`
	Name string `json:"name"`
	Year int32 `json:"year,omitempty"`
	Capacity Capacity `json:"capacity,omitempty"`
	Material string `json:"material,omitempty"`
	Categories []string `json:"categories,omitempty"`
	Version int32 `json:"version"`
}

func ValidateBlender(v *validator.Validator, blender *Blender) {
	v.Check(blender.Name != "", "name", "must be provided")
	v.Check(len(blender.Name) <= 500, "name", "must not be more than 500 bytes long")

	v.Check(blender.Year != 0, "year", "must be provided")
	v.Check(blender.Year >= 1970, "year", "must be greater than 1970")
	v.Check(blender.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(blender.Capacity != 0, "capacity", "must be provided")
	v.Check(blender.Capacity > 0, "capacity", "must be a positive integer")

	v.Check(blender.Categories != nil, "categories", "must be provided")
	v.Check(len(blender.Categories) >= 1, "categories", "must contain at least 1 category")
	v.Check(len(blender.Categories) <= 5, "categories", "must not contain more than 5 categories")
	v.Check(validator.Unique(blender.Categories), "categories", "must not contain duplicate values")
}
	