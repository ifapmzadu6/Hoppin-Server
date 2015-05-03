package action

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World.")

	decoder := json.NewDecoder(r.Body)
	var object Actions
	err := decoder.Decode(&object)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println(object)
}
