package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	templates *template.Template
)

func indexGetHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Hello")
	templates.ExecuteTemplate(writer, "index2.html", nil)
}
func main() {
	router := mux.NewRouter()
	templates = template.Must(template.ParseGlob("templates/[a-z]*.html"))
	router.HandleFunc("/", indexGetHandler).Methods("GET")
	http.Handle("/", router)
	http.ListenAndServe(":9090", nil)

}
