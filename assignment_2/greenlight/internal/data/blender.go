package data

import (
	"time"
	"database/sql" 
	"errors"

	"greenlight.alexedwards.net/internal/validator" 

	"github.com/lib/pq"
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


type BlenderModel struct {
	DB *sql.DB
}

func (m BlenderModel) Get(id int64) (*Blender, error) {
	if id < 1 {
			return nil, ErrRecordNotFound
		}

	query := `
	SELECT id, created_at, name, year, capacity, categories, version
	FROM blenders
	WHERE id = $1`
	
	var blender Blender

	err := m.DB.QueryRow(query, id).Scan(
		&blender.ID,
		&blender.CreatedAt,
		&blender.Name,
		&blender.Year,
		&blender.Capacity,
		pq.Array(&blender.Categories),
		&blender.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &blender, nil
}

func (m BlenderModel) Update(blender *Blender) error {
	query := `
		UPDATE blenders
		SET name = $1, year = $2, capacity = $3, categories = $4, version = version + 1
		WHERE id = $5
		RETURNING version`

	args := []interface{}{
		blender.Name,
		blender.Year,
		blender.Capacity,
		pq.Array(blender.Categories),
		blender.ID,
	}

	return m.DB.QueryRow(query, args...).Scan(&blender.Version)
}

func (m BlenderModel) Delete(id int64) error {
	// Return an ErrRecordNotFound error if the blender ID is less than 1.
	if id < 1 {
		return ErrRecordNotFound
	}

	// Construct the SQL query to delete the record.
	query := `
		DELETE FROM blenders
		WHERE id = $1`

	// Execute the SQL query using the Exec() method, passing in the id variable as
	// the value for the placeholder parameter. The Exec() method returns a sql.Result
	// object.
	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}

	// Call the RowsAffected() method on the sql.Result object to get the number of rows
	// affected by the query.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// If no rows were affected, we know that the blenders table didn't contain a record
	// with the provided ID at the moment we tried to delete it. In that case we
	// return an ErrRecordNotFound error.
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil	
}

func (m BlenderModel) Insert(blender *Blender) error {
	query := `
	INSERT INTO blenders (name, year, capacity, categories)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at, version`

	args := []interface{}{blender.Name, blender.Year, blender.Capacity, pq.Array(blender.Categories)}

	return m.DB.QueryRow(query, args...).Scan(&blender.ID, &blender.CreatedAt, &blender.Version)
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
