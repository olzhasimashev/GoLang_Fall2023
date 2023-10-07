package main

import (
	"fmt"
	"net/http"
	"time" // New import

	"greenlight.alexedwards.net/internal/data" // New import
	"greenlight.alexedwards.net/internal/validator" // New import
)

// Add a createBlenderHandler for the "POST /v1/blender" endpoint. For now we simply
// return a plain-text placeholder response.
func (app *application) createBlenderHandler(w http.ResponseWriter, r *http.Request) {
	// Declare an anonymous struct to hold the information that we expect to be in the
	// HTTP request body (note that the field names and types in the struct are a subset
	// of the Blender struct that we created earlier). This struct will be our *target
	// decode destination*.
	var input struct {
		Name string `json:"name"`
		Year int32 `json:"year"`
		Capacity data.Capacity `json:"capacity"`
		Categories []string `json:"categories"`
	}

	// Use the new readJSON() helper to decode the request body into the input struct.
	// If this returns an error we send the client the error message along with a 400
	// Bad Request status code, just like before.
	err := app.readJSON(w, r, &input)
	if err != nil {
		// Use the new badRequestResponse() helper.
		app.badRequestResponse(w, r, err)
		return
	}

	// Copy the values from the input struct to a new Blender struct.
	blender := &data.Blender{
		Name: input.Name,
		Year: input.Year,
		Capacity: input.Capacity,
		Categories: input.Categories,
	}

	// Initialize a new Validator.
	v := validator.New()

	// Call the ValidateBlender() function and return a response containing the errors if
	// any of the checks fail.
	if data.ValidateBlender(v, blender); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	
	fmt.Fprintf(w, "%+v\n", input)
}

// Add a showBlenderHandler for the "GET /v1/blender/:id" endpoint. For now, we retrieve
// the interpolated "id" parameter from the current URL and include it in a placeholder
// response.
func (app *application) showBlenderHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		// Use the new notFoundResponse() helper.
		app.notFoundResponse(w, r)
		return
	}

	// Create a new instance of the Blender struct, containing the ID we extracted from
	// the URL and some dummy data. Also notice that we deliberately haven't set a
	// value for the Year field.
	blender := data.Blender{
		ID: id,
		CreatedAt: time.Now(),
		Name: "Blender Experia 3000",
		Material: "Alluminium",
		Capacity: 3,
		Categories: []string{"blender", "cooking applience", "electronics"},
		Version: 1,
	}

	// Create an envelope{"blender": blender} instance and pass it to writeJSON(), instead
	// of passing the plain blender struct.
	err = app.writeJSON(w, http.StatusOK, envelope{"blender": blender}, nil)
	if err != nil {
		// Use the new serverErrorResponse() helper.
		app.serverErrorResponse(w, r, err)
	}	
}
