package main

import (
	"log"
	"net/http"
	"os"

	"./actions"
	"./users"
)

func main() {

	// Log
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Routing
	http.HandleFunc("/actions", actions.Handler)
	http.HandleFunc("/users/signup", users.SignUpHandler)

	// Listen
	if os.Getenv("DEBUG") == "1" {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := http.ListenAndServeTLS(":8080", "/ssl/cert.pem", "/ssl/key.pem", nil)
		if err != nil {
			log.Fatal(err)
		}
	}

}
