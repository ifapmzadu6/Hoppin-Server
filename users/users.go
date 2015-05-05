package users

import (
	"fmt"
	"net/http"

	"../mysql"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	password := "password"

	id, err := mysql.InsertUser(password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "{\"user\":"+Itoa(id)+", \"password\":"+password+"}")
	w.Header().Set("Content-Type", "application/json")
}
