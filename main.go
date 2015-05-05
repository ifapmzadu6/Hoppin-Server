package main

import (
	"log"
	"net/http"
	"os"

	"./actions"
	"./index"
	"./memcache"
	"./mysql"
	"./users"
)

func main() {
	err := mysql.Open()
	if err != nil {
		panic(err)
	}
	at := mysql.ActionType{String: "twitter"}
	a := mysql.Action{VideoId: "VideoId", Type: at, Time: 100, Start: 20, End: 30}
	mysql.InsertAction(a)
	defer mysql.Close()

	mc := memcache.Open()
	memcache.Set("value", "key", mc)
	//v, _ := memcache.Get("key", mc)

	http.HandleFunc("/", index.Handler)
	http.HandleFunc("/actions", actions.Handler)
	http.HandleFunc("/users/signup", users.SignUpHandler)

	if os.Getenv("DEBUG") == "1" {
		err = http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = http.ListenAndServeTLS(":8080", "/ssl/cert.pem", "/ssl/key.pem", nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
