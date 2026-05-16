package main

import (
	"fmt"
	"lense-code/controllers"
	"lense-code/templates"
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
	//tpl := views.Must(views.Parse("home.gohtml"))
	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))
	r.Get("/", controllers.Home(tpl))

	tplContact := views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tplContact))

	tplSignUp := views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/sign-up", controllers.StaticHandler(tplSignUp))

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}

}

//watch till section 51 intsall tailwind
