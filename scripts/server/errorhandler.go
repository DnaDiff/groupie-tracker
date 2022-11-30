package server

import (
	"net/http"
	"text/template"
)

type Error struct {
	Code    int
	Message string
}

var tmpl *template.Template

func httpError(w http.ResponseWriter, r *http.Request, code int) {
	err := Error{
		Code:	code,
		Message: http.StatusText(code),
	}
	errorHandler(w, r, &err)
}

func errorHandler(w http.ResponseWriter, r *http.Request, err *Error) {
	tmpl = template.Must(template.ParseFiles("templates/error.html"))
	w.WriteHeader(err.Code)
	tmpl.ExecuteTemplate(w, "error.html", err)
}
