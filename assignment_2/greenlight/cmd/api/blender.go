package main

import (
	"fmt"
	"net/http"
	"errors"

	"greenlight.alexedwards.net/internal/data"
	"greenlight.alexedwards.net/internal/validator"
)

func (app *application) createBlenderHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
		Year int32 `json:"year"`
		Capacity data.Capacity `json:"capacity"`
		Categories []string `json:"categories"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	blender := &data.Blender{
		Name: input.Name,
		Year: input.Year,
		Capacity: input.Capacity,
		Categories: input.Categories,
	}

	v := validator.New()

	if data.ValidateBlender(v, blender); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	
	err = app.models.Blenders.Insert(blender)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/blenders/%d", blender.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"blender": blender}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showBlenderHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	blender, err := app.models.Blenders.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"blender": blender}, nil)
	if err != nil {
	app.serverErrorResponse(w, r, err)
	}	
}

func (app *application) updateBlenderHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	blender, err := app.models.Blenders.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Name *string `json:"name"`
		Year *int32 `json:"year"`
		Capacity *data.Capacity `json:"capacity"`
		Categories []string `json:"categories"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Name != nil {
		blender.Name = *input.Name
	}
	if input.Year != nil {
		blender.Year = *input.Year
	}
	if input.Capacity != nil {
		blender.Capacity = *input.Capacity
	}
	if input.Categories != nil {
		blender.Categories = input.Categories
	}

	v := validator.New()

	if data.ValidateBlender(v, blender); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Blenders.Update(blender)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	
	err = app.writeJSON(w, http.StatusOK, envelope{"blender": blender}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteBlenderHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Blenders.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "blender successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
	
