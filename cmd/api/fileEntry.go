// Filename: cmd/api/fileEntry.go

package main

import (
	"fmt"
	"net/http"
	"quiz2/jamesfaber.net/internal/data"
	"quiz2/jamesfaber.net/internal/validator"
)

// createFileEntryHandler for the "POST /v1/Entryinfo" endpoint
func (app *application) createFileEntryHandler(w http.ResponseWriter, r *http.Request) {
	// Our target decode destination
	var input struct {
		Name    string   `json:"name"`
		Level   string   `json:"level"`
		Contact string   `json:"contact"`
		Phone   string   `json:"phone"`
		Email   string   `json:"email"`
		Website string   `json:"website"`
		Address string   `json:"address"`
		Mode    []string `json:"mode"`
	}
	//Initialize a new json.Decoder instance
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.readJSON(w, r, &input)
		if err != nil {
			app.badRequestResponse(w, r, err)
			return
		}
	}

	//Copy the values from the input struct to a new Entry struct
	fileEntry := &data.FileEntry{
		Name:    input.Name,
		Level:   input.Level,
		Contact: input.Contact,
		Phone:   input.Phone,
		Email:   input.Email,
		Website: input.Website,
		Address: input.Address,
		Mode:    input.Mode,
	}

	//Initialize a new Validator instance
	v := validator.New()
	//check the map to determine if there were any validation errors
	if data.ValidateFileEntry(v, fileEntry); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	//Display the request
	fmt.Fprintf(w, "%+v\n", input)
}

// // showRandomizeStringHandler for the "GET /vq.schools/:id" endpoint
func (app *application) showRandomizeStringHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	int1 := int(id)
	tools := &data.Tools{
		Int: int1,
	}
	// 	v := validator.New()
	// 	if data.ValidateInt(v, tools); !v.Valid() {
	// 		app.failedValidationResponse(w, r, v.Errors)
	// 		return
	// }
	strw := tools.GenerateRandomString(int1)
	data := envelope{
		"id":            int1,
		"random_string": strw,
	}
	err = app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
