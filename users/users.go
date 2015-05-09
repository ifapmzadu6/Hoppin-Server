package users

import (
	"crypto/rand"
	"encoding/base32"
	"io"
	"net/http"
	"strconv"
	"strings"

	"../mysql"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
	}

	device := r.URL.Query().Get("device")
	os := r.URL.Query().Get("os")
	udid, err := mysql.SelectUserDevice(device, os)
	if err != nil {
		nudid, err := mysql.InsertUserDevice(device, os)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		udid = nudid
	}

	b := make([]byte, 16)
	_, errt := io.ReadFull(rand.Reader, b)
	if errt != nil {
		http.Error(w, errt.Error(), http.StatusInternalServerError)
		return
	}
	password := strings.TrimRight(base32.StdEncoding.EncodeToString(b), "=")

	id, erriu := mysql.InsertUser(password, udid)
	if erriu != nil {
		http.Error(w, erriu.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"user\":" + strconv.FormatInt(id, 10) + ", \"password\":\"" + password + "\"}\n"))
}
