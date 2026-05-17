package controllers

import (
	"fmt"
	"lense-code/templates"
	"lense-code/views"
	"log"
	"net/http"
)

type Users struct {
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	tpl := views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	tpl.ExecuteTemplate(w, nil)
}
func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	log.Println("create new user")

	fmt.Fprintf(w, r.FormValue("email"))
	fmt.Fprintf(w, r.FormValue("password"))
}
