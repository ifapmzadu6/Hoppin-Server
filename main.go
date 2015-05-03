package main

import (
	"log"
	"net/http"

	"./action"
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

	v, _ := memcache.Get("key", mc)
	log.Println(v)

	http.HandleFunc("/actions/", action.Handler)
	http.ListenAndServe(":8080", nil)
}
