//Filename: cmd/api/healthcheck.go

package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	//create a map to hold ouor health check data
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	//convert our map into a JSON object
	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "the server encountered a problem and could not process yoour request.", http.StatusInternalServerError)
		return
	}

	//add a newline to make viewing on the terminal easier
	js = append(js, '\n')

	//specify that we will serve our responses using json
	w.Header().Set("Content-Type", "application/json")
	//write the []byte slice containing the JsoN response body
	w.Write(js)
}
