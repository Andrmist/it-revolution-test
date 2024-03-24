package utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type ErrorResponse struct {
	Error       string `json:"error,omitempty"`
	Description string `json:"description,omitempty"`
}

func SendError(w http.ResponseWriter, status int, name string, err error) {
	response := ErrorResponse{
		Error:       name,
		Description: err.Error(),
	}
	w.WriteHeader(status)
	if encErr := json.NewEncoder(w).Encode(&response); encErr != nil {
		SendError(w, http.StatusInternalServerError, "Internal Server Error", encErr)
	}
}

func SendValidationErrorResponse(w http.ResponseWriter, err validator.ValidationErrors) {
	w.WriteHeader(http.StatusBadRequest)
	response := ErrorResponse{
		Error:       "Bad Request",
		Description: "",
	}
	for _, e := range err {
		response.Description += fmt.Sprintf("field %s: %s\n", e.Field(), e.Error())
	}
	if encErr := json.NewEncoder(w).Encode(&response); encErr != nil {
		SendError(w, http.StatusInternalServerError, "Internal Server Error", encErr)
	}
}
