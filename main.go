package main

import (
	"log"
	"net/http"
	"os"

	"./actions"
	"./index"
	"./memcache"
	"./mysql"
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
	http.HandleFunc("/actions/", actions.Handler)

	if os.Getenv("DEBUG") == "1" {
		err = http.ListenAndServeTLS(":8080", "/ssl/cert.pem", "/ssl/key.pem", nil)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		if Exists("/ssl/cert.pem") == false {
			panic("/ssl/cert.pem is not exists.")
		}
		if Exists("/ssl/key.pem") == false {
			panic("/ssl/key.pem is not exists.")
		}

		log.Println("Start ListenAndServe :8080")
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
