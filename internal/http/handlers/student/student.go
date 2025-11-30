package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/AmarjitKaranSharma/golang-student-api/internal/storage"
	"github.com/AmarjitKaranSharma/golang-student-api/internal/types"
	"github.com/AmarjitKaranSharma/golang-student-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("Welcome to Students Api"))
		// log.Println(("Api"))

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		fmt.Println(err)

		// Specific error handling for empty body
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body"), http.StatusBadRequest))
			return
		}

		// Common error handling
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err, http.StatusBadRequest))
			return
		}

		// Validation error handling
		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		slog.Info("Creating a Student")

		lastId, err := storage.CreateStudent(student.Name, student.Email, student.Age)

		slog.Info("User Created", slog.String("userID", fmt.Sprintln(lastId)))

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
		}

		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}

func GetStudentById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("Student Id", r.PathValue("id"))

		var student types.Student
		id := r.PathValue("id")

		intId, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err, http.StatusInternalServerError))
			return
		}

		student, err = storage.GetStudentById(intId)

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err, http.StatusInternalServerError))
			return
		}

		response.WriteJson(w, http.StatusOK, student)
	}
}
