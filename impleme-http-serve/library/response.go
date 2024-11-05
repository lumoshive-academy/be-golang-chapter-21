package library

import (
	"be-golang-chapter-21/impleme-http-serve/model"
	"encoding/json"
	"net/http"
)

func BadResponse(w http.ResponseWriter, message string) {
	badResponse := model.Response{
		StatusCode: http.StatusBadRequest,
		Message:    "Error server",
		Data:       nil,
	}
	json.NewEncoder(w).Encode(badResponse)
}

func SuccessResponse(w http.ResponseWriter, message string, data any) {
	badResponse := model.Response{
		StatusCode: http.StatusBadRequest,
		Message:    "Error server",
		Data:       nil,
	}
	json.NewEncoder(w).Encode(badResponse)
}
