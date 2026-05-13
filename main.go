package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello Jafar!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}

//watch till section 11 finished
