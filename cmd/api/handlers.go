package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
}

func (app *application) patients(w http.ResponseWriter, r *http.Request) {

	patientList := []patient{{"Silvio Santos", "Hospital Legal"}, {"Hebe Camargo", "Hospital SP"}}

	js, _ := json.Marshal(patientList)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
