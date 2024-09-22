package main

import (
	"fmt"
	"html/template"
	"museum-server/api"
	"museum-server/data"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from a go program"))
}

func handleTempate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("./templates/index.tmpl")
	if err != nil {
		w.Write([]byte("Internal server error"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	html.Execute(w, data.GetAll())
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", handleHello)
	server.HandleFunc("/template", handleTempate)
	server.HandleFunc("/api", api.Get)
	server.HandleFunc("/api", api.Post)

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":3333", server)
	if err != nil {
		fmt.Errorf("error opening the server")
	}
}
