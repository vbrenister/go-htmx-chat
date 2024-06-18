package main

import (
	"fmt"
	"net/http"
)

func (app *application) handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
