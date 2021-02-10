package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
}

func (app *application) patients(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"patients": [{"name": "Silvio Santos", "hospital" : "Hospital Telesena"},{"name": "Hebe Camargo", "hospital" : "Hospital Gracinha"}]}`))
}
