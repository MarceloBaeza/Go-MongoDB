package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/mbaezahuenupil/go-mongodb-test/src/interfaces"
	"github.com/mbaezahuenupil/go-mongodb-test/src/models"
)

var once sync.Once
var instance *UserController

type UserController struct {
	UserService interfaces.UserServiceInterface
}

//NewUserController
func NewUserController(userService interfaces.UserServiceInterface) interfaces.UserControllerInterface {
	once.Do(func() {
		instance = &UserController{
			UserService: userService,
		}
	})
	return instance
}

func (ctrl *UserController) AddUser(w http.ResponseWriter, r *http.Request) {
	newUser := models.UserNew{}
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
	}
	validate := validator.New()
	err = validate.Struct(newUser)
	if err != nil {
		validationsErrors := err.(validator.ValidationErrors)
		responseError(validationsErrors, w, http.StatusBadRequest)
		return
	}
	data, _ := json.Marshal(newUser)
	fmt.Println("Data ", string(data))
	err = ctrl.UserService.Create(newUser)
	if err != nil {
		responseError(err, w, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (ctrl *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idUser := vars["id"]
	if idUser == "" {
		responseError(errors.New("id empty"), w, http.StatusBadRequest)
		return
	}
	err := ctrl.UserService.Delete(idUser)
	if err != nil {
		responseError(err, w, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
func (ctrl *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	usersList, err := ctrl.UserService.Read()
	if err != nil {
		responseError(err, w, http.StatusInternalServerError)
		return
	}
	responseBody := map[string]models.Users{"users": usersList}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(responseBody)
}
func (ctrl *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idUser := vars["id"]
	if idUser == "" {
		responseError(errors.New("id empty"), w, http.StatusBadRequest)
		return
	}
	updateUser := models.User{}
	err := json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		responseError(fmt.Errorf("error body parser %v", err.Error()), w, http.StatusBadRequest)
		return
	}
	err = ctrl.UserService.Update(updateUser, idUser)
	if err != nil {
		responseError(err, w, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func responseError(err error, w http.ResponseWriter, statusCode int) {
	validationsErrors := err.(validator.ValidationErrors)
	w.Header().Set("Content-Type", "application/json")
	responseBody := map[string]string{"error": validationsErrors.Error()}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(responseBody)

}
