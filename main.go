package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "Killers of the Flower Moon", Director: "Martin Scorsese"},
			},
		}

		tmpl.Execute(w, films)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second)

		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		newTmpl := template.Must(template.ParseFiles("index.html"))
		newTmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	}

	http.HandleFunc("/add-film/", h2)

	http.HandleFunc("/", h1)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
