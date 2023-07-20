package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, payload any) {
	res, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error marshaling response: %v", err)))
	}

	w.WriteHeader(status)
	w.Write(res)
}
