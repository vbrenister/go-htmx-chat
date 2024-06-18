package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func (app *application) render(w http.ResponseWriter, name string) {
	t, ok := app.templateCache[fmt.Sprintf("%s.html", name)]
	if !ok {
		panic("template not found " + name)
	}

	buf := new(bytes.Buffer)
	err := t.ExecuteTemplate(buf, "base", nil)
	if err != nil {
		panic(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	buf.WriteTo(w)
}
