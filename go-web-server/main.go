package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/tunnel", func(w http.ResponseWriter, r *http.Request) {
		for name, values := range r.Header {
			for _, value := range values {
				fmt.Fprintf(w, "%v: %v\n", name, value)
			}
		}
	})

	http.ListenAndServe(":8080", nil)
}
