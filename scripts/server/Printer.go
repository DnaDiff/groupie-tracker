package server

import (
	"log"
	"net/http"
	"text/template"

	get "groupie-tracker/scripts/data"
)

// type Everyone struct {
// 	Id           int
// 	Image        string
// 	Name         string
// 	Members      []string
// 	CreationDate int
// 	FirstAlbum   string
// 	Locations    []string
// 	ConcertDates []string
// 	Relations    string
// }

func Printer(w http.ResponseWriter, r *http.Request) {
	data := get.Epicmerge()
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.ExecuteTemplate(w, "index.html", data)
}
