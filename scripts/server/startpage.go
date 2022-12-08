package server

import (
	"groupie-tracker/scripts/data"
	"groupie-tracker/scripts/tools"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func isNumber(URL string) bool {
	nbr, err := strconv.Atoi(URL)
	if err != nil {
		return false
	}
	if nbr > 0 && nbr < 53 {
		return true
	}
	return false
}

func StartPage(w http.ResponseWriter, r *http.Request) {

	Artists := data.Compile()

	switch {
	case strings.HasPrefix(r.URL.Path, "/artist/") && isNumber(r.URL.Path[8:]):
		tmpl, err := template.ParseFiles("templates/artists.html")
		if err != nil {
			httpError(w, r, http.StatusInternalServerError)
			return
		}
		err = tmpl.ExecuteTemplate(w, "artists.html", Artists[tools.Atoi(r.URL.Path[8:])-1])
		if err != nil {
			httpError(w, r, http.StatusInternalServerError)
			return
		}
	case r.URL.Path == "/useless":
		tmpl, err := template.ParseFiles("templates/useless.html")
		if err != nil {
			httpError(w, r, http.StatusInternalServerError)
			return
		}
		err = tmpl.ExecuteTemplate(w, "useless.html", Artists)
		if err != nil {
			httpError(w, r, http.StatusInternalServerError)
			return
		}
	case r.URL.Path == "/":
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			httpError(w, r, http.StatusInternalServerError)
			return
		}
		err = tmpl.ExecuteTemplate(w, "index.html", Artists)
		if err != nil {
			httpError(w, r, http.StatusInternalServerError)
			return
		}
	default:
		httpError(w, r, http.StatusNotFound)
		return
	}
}
