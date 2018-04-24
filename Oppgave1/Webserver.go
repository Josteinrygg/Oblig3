package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", response)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func response(w http.ResponseWriter, r *http.Request) {
	msg := "Hello, client"
	w.Write([]byte(msg))
}