package main

import (
	"log"
	"net/http"

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

	log.Println("Start ListenAndServe :8080")
	http.ListenAndServe(":8080", nil)
}
