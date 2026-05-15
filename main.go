package main

import (
	"fmt"
	"lense-code/controllers"
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

func main() {
	r := chi.NewRouter()
	tpl := views.Must(views.Parse("home.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tplContact := views.Must(views.Parse("contact.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tplContact))

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}

}

//watch till section 44 finished
