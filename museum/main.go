package main

import (
	"fmt"
	"html/template"
	"net/http"

	"fe.com/go/museum/api"
	"fe.com/go/museum/public/data"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from a Go program"))
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	html.Execute(w, data.GetAll())
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", handleHello)
	server.HandleFunc("/template", handleTemplate)
	server.HandleFunc("/api/exhibitions", api.Get)

	fileSys := http.FileServer(http.Dir("./public/"))
	server.Handle("/", fileSys)

	err := http.ListenAndServe(":3333", server)
	if err == nil {
		fmt.Println("Error while opening the server")
	}
}
