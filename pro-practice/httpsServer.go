package main

import (
	"fmt"
	"net/http"
)

var PORT = ":1443"

func main() {
	http.HandleFunc("/", Default)
	fmt.Println("Listening to port number", PORT)

	err := http.ListenAndServeTLS(PORT, "servercrt", "server.key", nil)
	if err != nil {
		fmt.Println("Listen And Serve TLS: ", err)
		return
	}

}

func Default(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is an example HTTPS server!")
}
