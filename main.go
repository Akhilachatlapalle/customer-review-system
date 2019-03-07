package main

import (
	"fmt"
	"net/http"
)

func main() {
	ApplicationPort := "8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "My Go Application!")
	})
	http.ListenAndServe(":"+ApplicationPort, nil)
}
