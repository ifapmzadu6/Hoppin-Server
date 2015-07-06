package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	"./actions"
	"./histograms"
	"./mysql"
	"./users"
)

func main() {

	// Log
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// MySQL
	if err := setupMySQL(); err != nil {
		log.Fatal(err.Error())
		return
	}

	// Routing
	http.HandleFunc("/actions", actions.Handler)
	http.HandleFunc("/users/signup", users.SignUpHandler)
	http.HandleFunc("/videos/histogram", histograms.Handler)

	// Listen
	if os.Getenv("DEBUG") == "1" {
		err := http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
		if err != nil {
			log.Fatal(err.Error())
		}
	} else {
		err := http.ListenAndServeTLS(":8080", "/ssl/cert.pem", "/ssl/key.pem", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
		if err != nil {
			log.Fatal(err.Error())
		}
	}

}

func setupMySQL() error {
	db, dbErr := mysql.Connect()
	if dbErr != nil {
		return dbErr
	}
	defer mysql.Close(db)

	if err := mysql.CreateDataBase(db); err != nil {
		return err
	}
	if err := mysql.SelectDataBase(db); err != nil {
		return err
	}
	if err := mysql.CreateTables(db); err != nil {
		return err
	}
	return nil
}
