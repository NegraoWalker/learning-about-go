package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type Usuario struct {
	Nome  string
	Email string
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		usuario := Usuario{
			"Walker",
			"walker@email.com",
		}
		templates.ExecuteTemplate(w, "index.html", usuario)
	})
	log.Fatal(http.ListenAndServe(":5000", nil))
	fmt.Println("Executando na porta 5000...")
}
