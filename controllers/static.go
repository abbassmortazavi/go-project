package controllers

import (
	"lense-code/views"
	"net/http"
)

func StaticHandler(tmpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, nil)
	}
}

func Home(tmpl views.Template) http.HandlerFunc {
	data := []struct {
		Title   string
		Content string
	}{
		{Title: "Index", Content: "Welcome"},
		{Title: "Login", Content: "Login"},
		{Title: "Logout", Content: "Logout"},
		{Title: "Error", Content: "Error"},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, data)
	}
}
