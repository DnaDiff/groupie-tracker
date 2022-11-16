package server

import (
	"net/http"
)

type Error struct {
	Code    int
	Message string
}

// handles 500 errors
func error500(w http.ResponseWriter, r *http.Request) {
	d := Error{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
	}
	errorHandler(w, r, &d)
}

func errorHandler(w http.ResponseWriter, r *http.Request, d *Error) {
	// tmpl = template.Must(template.ParseFiles("templates/error.html"))
	w.WriteHeader(d.Code)
	// tmpl.ExecuteTemplate(w, "error.html", d)
}
