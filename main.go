package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Hello Jafar!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1>")
}

//func pathHandler(w http.ResponseWriter, r *http.Request) {
//	switch r.URL.Path {
//	case "/":
//		homeHandler(w, r)
//	case "/contact":
//		contactHandler(w, r)
//	default:
//		http.Error(w, "404 not found.", http.StatusNotFound)
//	}
//}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "404 not found.", http.StatusNotFound)
	}
}
func main() {
	var router Router
	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}

}

//watch till section 15 finished
