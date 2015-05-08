package users

import (
	"crypto/rand"
	"encoding/base32"
	"io"
	"net/http"
	"strconv"

	"../mysql"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
	}

	b := make([]byte, 32)
	_, errt := io.ReadFull(rand.Reader, b)
	if errt != nil {
		http.Error(w, errt.Error(), http.StatusInternalServerError)
		return
	}
	password := base32.StdEncoding.EncodeToString(b)

	id, err := mysql.InsertUser(password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"user\":" + strconv.FormatInt(id, 10) + ", \"password\":\"" + password + "\"}\n"))
}
