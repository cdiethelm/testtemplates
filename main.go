package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	startHTTP()
}

type PageData struct {
	PageTitle string
	Options   []string
}

func startHTTP() {
	r := mux.NewRouter()

	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "SecureKey"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"POST", "GET"})

	r.HandleFunc("/test/", testHandler).Methods("GET")
	r.HandleFunc("/template/", templateHandler).Methods("GET")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(r)))
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Chris is great</h1>")
}

func templateHandler(w http.ResponseWriter, r *http.Request) {

	var data PageData
	data = PageData{
		PageTitle: "Chrises Space",
		Options:   []string{"hyundai", "volvo", "tesla"},
	}

	tmpl := template.Must(template.ParseFiles("templates/test.html"))

	tmpl.Execute(w, data)

}
