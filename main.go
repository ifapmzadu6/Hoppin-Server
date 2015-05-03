package main

import (
	"log"
	"net/http"

	"./action"
	"./memcache"
	"./mysql"
)

/*
 15 type Action struct {
	  16     Id      int
	   17     VideoId string
	    18     Type    ActionType
		 19     Time    int
		  20     Start   int
		   21     End     int
		    22 }
			 23
			  24 type ActionType struct {
				   25     Id     int
				    26     String string
					 27 }
*/

func main() {
	db, _ := mysql.Open()
	at := mysql.ActionType{String: "twitter"}
	a := mysql.Action{VideoId: "VideoId", Type: at, Time: 100, Start: 20, End: 30}
	mysql.InsertAction(db, a)
	defer mysql.Close(db)

	mc := memcache.Open()
	memcache.Set("value", "key", mc)

	v, _ := memcache.Get("key", mc)
	log.Println(v)

	http.HandleFunc("/action/", action.Handler)
	http.ListenAndServe(":8080", nil)
}
