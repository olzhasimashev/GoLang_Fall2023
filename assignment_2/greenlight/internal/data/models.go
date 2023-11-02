package data

import (
	"database/sql"
	"errors"
)

// Define a custom ErrRecordNotFound error. We'll return this from our Get() method when
// looking up a movie that doesn't exist in our database.
var (
	ErrRecordNotFound = errors.New("record not found")
)

// Create a Models struct which wraps the BlenderModel. We'll add other models to this,
// like a UserModel and PermissionModel, as our build progresses.
type Models struct {
	Blenders BlenderModel
}

// For ease of use, we also add a New() method which returns a Models struct containing
// the initialized BlenderModel.
func NewModels(db *sql.DB) Models {
	return Models{
		Blenders: BlenderModel{DB: db},
	}
}