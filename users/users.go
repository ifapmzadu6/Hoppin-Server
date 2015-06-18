package users

import (
	"crypto/rand"
	"database/sql"
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
		return
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

	udid, udidErr := getUserDeviceId(db, device)
	if udidErr != nil {
		http.Error(w, udidErr.Error(), http.StatusInternalServerError)
		log.Println(udidErr.Error())
		return
	}

	uoid, uoidErr := getUserOSId(db, os)
	if uoidErr != nil {
		http.Error(w, uoidErr.Error(), http.StatusInternalServerError)
		log.Println(uoidErr.Error())
		return
	}

	password, pwErr := getRandomPassword()
	if pwErr != nil {
		http.Error(w, pwErr.Error(), http.StatusInternalServerError)
		log.Println(pwErr.Error())
		return
	}

	id, erriu := mysql.InsertUser(db, password, udid, uoid)
	if erriu != nil {
		http.Error(w, erriu.Error(), http.StatusInternalServerError)
		log.Println(erriu.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"user\":" + strconv.FormatInt(id, 10) + ", \"password\":\"" + password + "\"}\n"))
}

func getUserDeviceId(db *sql.DB, device string) (int64, error) {
	udid, udidErr := mysql.SelectUserDevice(db, device)
	if udidErr != nil {
		nudid, nudidErr := mysql.InsertUserDevice(db, device)
		if nudidErr != nil {
			return -1, nudidErr
		}
		return nudid, nil
	}
	return udid, nil
}

func getUserOSId(db *sql.DB, os string) (int64, error) {
	uoid, uoidErr := mysql.SelectUserOS(db, os)
	if uoidErr != nil {
		nuoid, nuoidErr := mysql.InsertUserOS(db, os)
		if nuoidErr != nil {
			return -1, nuoidErr
		}
		return nuoid, nil
	}
	return uoid, nil
}

func getRandomPassword() (string, error) {
	b := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return "", err
	}
	return strings.TrimRight(base32.StdEncoding.EncodeToString(b), "="), nil
}
