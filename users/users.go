package users

import (
	"crypto/rand"
	"encoding/base32"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"../mysql"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
	}

	db, dberr := mysql.Open()
	if dberr != nil {
		http.Error(w, dberr.Error(), http.StatusInternalServerError)
		log.Println(dberr.Error())
		return
	}
	defer mysql.Close(db)

	device := r.URL.Query().Get("device")
	os := r.URL.Query().Get("os")
	udid, err := mysql.SelectUserDevice(db, device, os)
	if err != nil {
		nudid, err := mysql.InsertUserDevice(db, device, os)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}
		udid = nudid
	}

	b := make([]byte, 16)
	_, errt := io.ReadFull(rand.Reader, b)
	if errt != nil {
		http.Error(w, errt.Error(), http.StatusInternalServerError)
		log.Println(errt.Error())
		return
	}
	password := strings.TrimRight(base32.StdEncoding.EncodeToString(b), "=")

	id, erriu := mysql.InsertUser(db, password, udid)
	if erriu != nil {
		http.Error(w, erriu.Error(), http.StatusInternalServerError)
		log.Println(erriu.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"user\":" + strconv.FormatInt(id, 10) + ", \"password\":\"" + password + "\"}\n"))
}
