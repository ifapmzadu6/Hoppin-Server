package actions

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"../mysql"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
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

	decoder := json.NewDecoder(r.Body)
	var as Actions
	if err := decoder.Decode(&as); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	for _, a := range as.Actions {
		actionId, aErr := getActionTypeId(db, a.Type)
		if aErr != nil {
			http.Error(w, aErr.Error(), http.StatusInternalServerError)
			log.Println(aErr.Error())
			return
		}

		videoId, vErr := getVideoId(db, a.VideoId)
		if vErr != nil {
			http.Error(w, vErr.Error(), http.StatusInternalServerError)
			log.Println(vErr.Error())
			return
		}

		if naErr := mysql.InsertAction(db, videoId, actionId, int64(userId), a.Time, a.Start, a.End); naErr != nil {
			http.Error(w, naErr.Error(), http.StatusInternalServerError)
			log.Println(naErr.Error())
			return
		}
	}

	return
}

func getVideoId(db *sql.DB, name string) (int64, error) {
	id, err := mysql.SelectVideoId(db, name)
	if err != nil {
		tid, tErr := getVideoTypeId(db, "YouTube")
		if tErr != nil {
			return -1, tErr
		}

		return mysql.InsertVideo(db, name, tid)
	}
	return id, nil
}

func getVideoTypeId(db *sql.DB, name string) (int64, error) {
	id, err := mysql.SelectVideoType(db, name)
	if err != nil {
		return mysql.InsertVideoType(db, name)
	}
	return id, nil
}

func getActionTypeId(db *sql.DB, name string) (int64, error) {
	id, err := mysql.SelectActionTypeByName(db, name)
	if err != nil {
		return mysql.InsertActionType(db, name)
	}
	return int64(id), nil
}
