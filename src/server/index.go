package server

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mbaezahuenupil/go-mongodb-test/src/controller"
	"github.com/mbaezahuenupil/go-mongodb-test/src/middleware"
	"github.com/mbaezahuenupil/go-mongodb-test/src/repositories/userrepository"
	userservice "github.com/mbaezahuenupil/go-mongodb-test/src/services/user.service"
)

func ConfigureServer(router *mux.Router) {
	userRepository := userrepository.NewUserRepository(os.Getenv("USER_COLLECTION"))
	userService := userservice.NewUserService(userRepository)
	crtl := controller.NewUserController(userService)
	router.Use(middleware.LoggingMiddleware)
	router.HandleFunc("/healthCheck", controller.HealthCheck).Methods(http.MethodGet)
	router.HandleFunc("/", crtl.AddUser).Methods(http.MethodPost)
	router.HandleFunc("/{id}", crtl.DeleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/getAllUsers", crtl.GetAllUsers).Methods(http.MethodGet)
	// Update
	subRouterUpdate := router.PathPrefix("/update").Subrouter()
	subRouterUpdate.HandleFunc("/{id}", crtl.UpdateUser).Methods(http.MethodPatch)

}
