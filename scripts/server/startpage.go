package server

import (
	get "groupie-tracker/scripts/Data"
	"html/template"
	"net/http"
)

func StartPage(w http.ResponseWriter, r *http.Request) {

	Artists := get.Epicmerge()

	switch {
	// case r.Method != "GET":
	// 	error404(w, r)
	// case r.URL.Path != "/":
	// 	error404(w, r)
	// case r.URL.Path != "/artists":
	default:

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			error500(w, r)
			return
		}
		err2 := tmpl.ExecuteTemplate(w, "index.html", Artists)
		if err2 != nil {
			error500(w, r)
			return
		}

	}
}
