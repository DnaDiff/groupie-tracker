package server

import (
	get "groupie-tracker/scripts/Data"
	"net/http"
	"text/template"
)

// func StartPage(w http.ResponseWriter, r *http.Request) {

// if r.URL.Path != "/" || r.Method != "GET" {
// 	error404(w, r)
// 	return
// } else {
// 	tmpl, err := template.ParseFiles("templates/index.html")
// 	if err != nil {
// 		error500(w, r)
// 		return
// 	}
// 	err2 := tmpl.ExecuteTemplate(w, "index.html", get.Epicmerge())
// 	if err2 != nil {
// 		error500(w, r)
// 		return
// 	}
// }
// }

func StartPage(w http.ResponseWriter, r *http.Request) {
	nice := get.Epicmerge()
	index, err := template.ParseFiles("templates/cocc.html")
	if err != nil {
		error500(w, r)
		return
	}
	tmpl = template.Must(index, err)

	tmpl.ExecuteTemplate(w, "index.html", nice)
}
