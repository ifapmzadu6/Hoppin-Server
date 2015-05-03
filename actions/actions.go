package actions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../mysql"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var as Actions
	err := decoder.Decode(&as)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i, a := range as.Actions {
		fmt.Print(fmt.Sprintf("index:%d,value:%d\n", i, a))

		nat := mysql.ActionType{String: a.Type}
		na := mysql.Action{VideoId: a.VideoId, Type: nat, Time: a.Time, Start: a.Start, End: a.End}
		mysql.InsertAction(na)
	}
}
