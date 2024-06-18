package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "hello world")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("listening at localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
