package main

import (
	"net/http"

	"./action"
	"./mysql"
)

func main() {
	db := mysql.Open()
	defer mysql.Close(db)

	http.HandleFunc("/", action.Handler)
	http.ListenAndServe(":8080", nil)
}
