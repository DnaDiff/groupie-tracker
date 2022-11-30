package server

import (
	get "groupie-tracker/scripts/Data"
	"html/template"
	"net/http"
	"strconv"
)

func CheckNumber(URL string) bool {
	nbr, err := strconv.Atoi(URL)
	if err != nil {
		return false
	}
	if nbr > 0 && nbr < 53 {
		return true
	}
	return false
}

func Atoi(s string) int {
	nice, _ := strconv.Atoi(s)
	return nice
}

func StartPage(w http.ResponseWriter, r *http.Request) {

	Artists := get.Epicmerge()
	// fmt.Println(Artists[Atoi(r.URL.Path[1:])].Id)

	switch {
	case CheckNumber(r.URL.Path[1:]):
		tmpl, err := template.ParseFiles("templates/artists.html")
		if err != nil {
			error500(w, r)
			return
		}
		err2 := tmpl.ExecuteTemplate(w, "artists.html", Artists[Atoi(r.URL.Path[1:])-1])
		if err2 != nil {
			error500(w, r)
			return
		}
	case r.URL.Path == "/":
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
	default:
		error404(w, r)
		return
	}

	// if CheckNumber(r.URL.Path[1:]) {
	// 	tmpl, err := template.ParseFiles("templates/artists.html")
	// 	if err != nil {
	// 		error500(w, r)
	// 		return
	// 	}
	// 	err2 := tmpl.ExecuteTemplate(w, "artists.html", Artists[Atoi(r.URL.Path[1:])-1])
	// 	if err2 != nil {
	// 		error500(w, r)
	// 		return
	// 	}
	// } else if r.URL.Path == "/" {
	// 	tmpl, err := template.ParseFiles("templates/index.html")
	// 	if err != nil {
	// 		error500(w, r)
	// 		return
	// 	}
	// 	err2 := tmpl.ExecuteTemplate(w, "index.html", Artists)
	// 	if err2 != nil {
	// 		error500(w, r)
	// 		return
	// 	} else if r.URL.Path != "/" {
	// 		error404(w, r)
	// 	}
	// }

	// switch {
	// // case r.Method != "GET":
	// // 	error404(w, r)
	// // case r.URL.Path != "/":
	// // 	error404(w, r)
	// // case r.URL.Path != "/artists":
	// default:

	// 	tmpl, err := template.ParseFiles("templates/index.html")
	// 	if err != nil {
	// 		error500(w, r)
	// 		return
	// 	}
	// 	err2 := tmpl.ExecuteTemplate(w, "index.html", Artists)
	// 	if err2 != nil {
	// 		error500(w, r)
	// 		return
	// 	}

	// }
}
