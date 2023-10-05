package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/iDominate/golang-postgres/models"
)

func ReturnResponse(w http.ResponseWriter, statusCode int, data interface{}, message string) {
	w.WriteHeader(statusCode)
	w.Header().Set("ContentType", "application/json")
	response := models.Response{
		StatusCode: int16(statusCode),
		Message:    message,
		Data:       data,
	}
	res, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%s\n", res)
}
