package main

import (
	"boke/control"
	"net/http"
)

func main() {
	http.Handle(`/static/`, http.StripPrefix(`/static/`, http.FileServer(http.Dir("static"))))
	http.HandleFunc(`/`, control.ViewIndex)

	http.ListenAndServe(`:80`, nil)
}
