package main

import (
	"html/template"
	"log"
	"net/http"
)

type MyHandler struct{}

func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, "Hello World!!")
}

func main() {
	http.HandleFunc("/top", top)
	http.ListenAndServe(":8080", nil)
}
