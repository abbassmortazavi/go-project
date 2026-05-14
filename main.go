package main

import (
	"fmt"
	"lense-code/views"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func render(w http.ResponseWriter, name string) {
	t, err := views.Parse(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "home.gohtml")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "contact.gohtml")
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}

}

//watch till section 29 finished
