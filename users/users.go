package users

import (
	"fmt"
	"net/http"
	"strconv"

	"../mysql"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	password := "password"

	id, err := mysql.InsertUser(password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "{\"user\":"+strconv.FormatInt(id, 10)+", \"password\":"+password+"}")
	w.Header().Set("Content-Type", "application/json")
}
