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
