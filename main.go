package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello world")

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":5000", nil)
}
