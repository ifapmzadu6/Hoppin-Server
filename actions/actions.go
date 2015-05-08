package actions

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../mysql"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
	}

	id := r.URL.Query().Get("user")
	password := r.URL.Query().Get("password")
	if err := mysql.ValidateUserByStr(id, password); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var as Actions
	err := decoder.Decode(&as)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, a := range as.Actions {
		var natnid int64
		var err error
		natid, err := mysql.SelectActionTypeByName(a.Type)
		if err != nil {
			tnatnid, natnerr := mysql.InsertActionType(a.Type)
			if natnerr != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			natnid = tnatnid
		} else {
			natnid = int64(natid)
		}

		nat := mysql.ActionType{Id: natnid, Name: a.Type}
		na := mysql.Action{VideoId: a.VideoId, Type: nat, Time: a.Time, Start: a.Start, End: a.End}
		userId, _ := strconv.Atoi(id)
		err = mysql.InsertAction(na, userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
