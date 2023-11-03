package data

import (
	"context"
	"database/sql" 
	"errors"
	"fmt"
	"time"

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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
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
		WHERE id = $5 AND version = $6
		RETURNING version`

	args := []interface{}{
		blender.Name,
		blender.Year,
		blender.Capacity,
		pq.Array(blender.Categories),
		blender.ID,
		blender.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&blender.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (m BlenderModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM blenders
		WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	return m.DB.QueryRowContext(ctx, query, args...).Scan(&blender.ID, &blender.CreatedAt, &blender.Version)
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

func (m BlenderModel) GetAll(name string, categories []string, filters Filters) ([]*Blender, Metadata, error) {
	query := fmt.Sprintf(`
		SELECT count(*) OVER(), id, created_at, name, year, capacity, categories, version
		FROM blenders
		WHERE (to_tsvector('simple', name) @@ plainto_tsquery('simple', $1) OR $1 = '')
		AND (categories @> $2 OR $2 = '{}')
		ORDER BY %s %s, id ASC
		LIMIT $3 OFFSET $4`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{name, pq.Array(categories), filters.limit(), filters.offset()}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	blenders := []*Blender{}

	for rows.Next() {
		var blender Blender

		err := rows.Scan(
			&totalRecords,
			&blender.ID,
			&blender.CreatedAt,
			&blender.Name,
			&blender.Year,
			&blender.Capacity,
			pq.Array(&blender.Categories),
			&blender.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}
		
		blenders = append(blenders, &blender)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return blenders, metadata, nil
}
	