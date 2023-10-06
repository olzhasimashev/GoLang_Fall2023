package main

import (
	"fmt"
	"net/http"
	"time" // New import

	"greenlight.alexedwards.net/internal/data" // New import
)

// Add a createMovieHandler for the "POST /v1/movies" endpoint. For now we simply
// return a plain-text placeholder response.
func (app *application) createCookingApplienceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new cooking applience")
}

// Add a showMovieHandler for the "GET /v1/movies/:id" endpoint. For now, we retrieve
// the interpolated "id" parameter from the current URL and include it in a placeholder
// response.
func (app *application) showCookingApplienceHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		// Use the new notFoundResponse() helper.
		app.notFoundResponse(w, r)
		return
	}

	// Create a new instance of the Movie struct, containing the ID we extracted from
	// the URL and some dummy data. Also notice that we deliberately haven't set a
	// value for the Year field.
	cooking_applience := data.CookingApplience{
		ID: id,
		CreatedAt: time.Now(),
		Name: "Some Blender and Mixer",
		Material: "Alluminium",
		Categories: []string{"blender", "mixer"},
		Version: 1,
	}

	// Create an envelope{"cooking_applience": cooking_applience} instance and pass it to writeJSON(), instead
	// of passing the plain cooking_applience struct.
	err = app.writeJSON(w, http.StatusOK, envelope{"cooking_applience": cooking_applience}, nil)
	if err != nil {
		// Use the new serverErrorResponse() helper.
		app.serverErrorResponse(w, r, err)
	}	
}
