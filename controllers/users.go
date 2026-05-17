package controllers

import (
	"lense-code/templates"
	"lense-code/views"
	"net/http"
)

type Users struct {
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	tpl := views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	tpl.ExecuteTemplate(w, nil)
}
