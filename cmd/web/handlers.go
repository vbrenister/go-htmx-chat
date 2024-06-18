package main

import (
	"net/http"
)

func (app *application) handleIndex(w http.ResponseWriter, r *http.Request) {
	app.render(w, "chat")
}
