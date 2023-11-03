package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict = errors.New("edit conflict")
)	

type Models struct {
	Blenders BlenderModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Blenders: BlenderModel{DB: db},
	}
}