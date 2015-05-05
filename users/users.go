package users

import (
	"net/http"
	"strconv"

	"../mysql"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	password := "password"

	id, err := mysql.InsertUser(password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"user\":" + strconv.FormatInt(id, 10) + ", \"password\":" + password + "}\n"))
}
