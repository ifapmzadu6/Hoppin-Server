package histograms

import (
	"database/sql"
	"log"
	"net/http"

	"../mysql"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.NotFound(w, r)
		log.Println(r)
		return
	}
	db, dberr := mysql.Open()
	if dberr != nil {
		http.Error(w, dberr.Error(), http.StatusUnauthorized)
		log.Println(dberr.Error())
		return
	}
	defer mysql.Close(db)

	id := r.URL.Query().Get("user")
	password := r.URL.Query().Get("password")

	userId, uErr := mysql.ValidateUserByStr(db, id, password)
	if uErr != nil {
		http.Error(w, uErr.Error(), http.StatusUnauthorized)
		log.Println(uErr.Error())
		return
	}

	log.Println(userId)
}

func getVideoId(db *sql.DB, name string) (int64, error) {
	id, err := mysql.SelectVideoId(db, name)
	if err != nil {
		return -1, err
	}
	return id, nil
}
