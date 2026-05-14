package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type Template struct {
	templates *template.Template
}

func Parse(filename string) (Template, error) {
	path := filepath.Join("templates", filename)
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return Template{}, fmt.Errorf("template parsing error: %w", err)
	}
	return Template{
		templates: tmpl,
	}, err
}

func (t *Template) ExecuteTemplate(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.templates.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
