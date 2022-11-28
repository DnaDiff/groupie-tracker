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

func error500(w http.ResponseWriter, r *http.Request) {
	d := Error{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
	}
	errorHandler(w, r, &d)
}

func error404(w http.ResponseWriter, r *http.Request) {
	d := Error{
		Code:    http.StatusNotFound,
		Message: "Page not found",
	}
	errorHandler(w, r, &d)
}

func error405(w http.ResponseWriter, r *http.Request) {
	d := Error{
		Code:    http.StatusMethodNotAllowed,
		Message: "Method not allowed",
	}
	errorHandler(w, r, &d)
}

func error400(w http.ResponseWriter, r *http.Request) {
	d := Error{
		Code:    http.StatusBadRequest,
		Message: "Bad request",
	}
	errorHandler(w, r, &d)
}

func errorHandler(w http.ResponseWriter, r *http.Request, d *Error) {
	tmpl = template.Must(template.ParseFiles("templates/error.html"))
	w.WriteHeader(d.Code)
	tmpl.ExecuteTemplate(w, "error.html", d)
}
