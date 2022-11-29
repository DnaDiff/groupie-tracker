package server

import (
	"fmt"
	"log"
	"net/http"
)

var Port = "80"

func StartServer() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/", StartPage)
	fmt.Println("Server started on port", Port)
	err := http.ListenAndServe(":"+Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
