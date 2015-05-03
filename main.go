package main

import (
	"log"
	"net/http"

	"./action"
	"./memcache"
	"./mysql"
)

func main() {
	db, _ := mysql.Open()
	defer mysql.Close(db)

	mc := memcache.Open()
	memcache.Set("value", "key", mc)

	v, _ := memcache.Get("key", mc)
	log.Println(v)

	http.HandleFunc("/action/", action.Handler)
	http.ListenAndServe(":8080", nil)
}
