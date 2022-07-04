package interfaces

import (
	"net/http"

	"github.com/mbaezahuenupil/go-mongodb-test/src/models"
)

//UserRepositoryInterface interface
type UserRepositoryInterface interface {
	Create(user models.User) error
	Read() (models.Users, error)
	Update(user models.User, userid string) error
	Delete(userid string) error
}

//UserServiceInterface interface
type UserServiceInterface interface {
	Create(userNew models.UserNew) error
	Read() (models.Users, error)
	Update(user models.User, userid string) error
	Delete(userid string) error
}

//UserControllerInterface interface
type UserControllerInterface interface {
	AddUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
}
