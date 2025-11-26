package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/AmarjitKaranSharma/golang-student-api/internal/types"
	"github.com/go-playground/validator/v10"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error, status int) types.Response {
	return types.Response{
		Status: status,
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) types.Response {
	var errMsg []string

	for _, err := range errs {
		switch err.Tag() {
		case "required":
			errMsg = append(errMsg, fmt.Sprintf("Field %s is required", err.Field()))
		default:
			errMsg = append(errMsg, fmt.Sprintf("Field %s is invalid", err.Field()))
		}
	}

	return types.Response{Status: http.StatusBadRequest, Error: strings.Join(errMsg, ",")}
}
