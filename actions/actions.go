package actions

import (
	"encoding/json"
	"net/http"

	"../mysql"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
	}

	decoder := json.NewDecoder(r.Body)
	var as Actions
	err := decoder.Decode(&as)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, a := range as.Actions {
		nat := mysql.ActionType{String: a.Type}
		na := mysql.Action{VideoId: a.VideoId, Type: nat, Time: a.Time, Start: a.Start, End: a.End}
		err = mysql.InsertAction(na)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
