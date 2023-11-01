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

	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		fmt.Println(err)
	}
}

/*

By using openssl, I created a self signed certificated for hte HTTPS server and it worked.
Therefore the files "server.crt" and "server.key".

*/
