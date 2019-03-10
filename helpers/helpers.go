package helpers

import (
	"customer-review-system/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func IsEmpty(data string) bool {
	if len(data) == 0 {
		return true
	} else {
		return false
	}
}

func SendResp(w http.ResponseWriter, resp models.Response, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Could not encode %+v to json. Got error: %s", resp, err.Error()))
	}
}
