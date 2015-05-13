package actions

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
	if err := mysql.ValidateUserByStr(db, id, password); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		log.Println(err.Error())
		return
	}

	decoder := json.NewDecoder(r.Body)
	var as Actions
	err := decoder.Decode(&as)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	for _, a := range as.Actions {
		var natnid int64
		var err error
		natid, err := mysql.SelectActionTypeByName(db, a.Type)
		if err != nil {
			tnatnid, natnerr := mysql.InsertActionType(db, a.Type)
			if natnerr != nil {
				http.Error(w, natnerr.Error(), http.StatusBadRequest)
				log.Println(natnerr.Error())
				return
			}
			natnid = tnatnid
		} else {
			natnid = int64(natid)
		}

		nat := mysql.ActionType{Id: natnid, Name: a.Type}
		na := mysql.Action{VideoId: a.VideoId, Type: nat, Time: a.Time, Start: a.Start, End: a.End}
		userId, _ := strconv.Atoi(id)
		err = mysql.InsertAction(db, na, userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}
	}
}
