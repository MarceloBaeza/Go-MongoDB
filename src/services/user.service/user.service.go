package userservice

import (
	"sync"
	"time"

	"github.com/mbaezahuenupil/go-mongodb-test/src/interfaces"
	"github.com/mbaezahuenupil/go-mongodb-test/src/models"
)

var instance *UserService
var once sync.Once

//UserService struct
type UserService struct {
	UserRepository interfaces.UserRepositoryInterface
}

//NewUserService constructor
func NewUserService(userRepository interfaces.UserRepositoryInterface) interfaces.UserServiceInterface {
	once.Do(func() {
		instance = &UserService{
			UserRepository: userRepository,
		}
	})
	return instance
}

func (us *UserService) Create(userNew models.UserNew) error {
	user := models.User{
		Name:      userNew.Name,
		Email:     userNew.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return us.UserRepository.Create(user)
}

func (us *UserService) Read() (models.Users, error) {
	users, err := us.UserRepository.Read()
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (us *UserService) Update(user models.User, userid string) error {
	return us.UserRepository.Update(user, userid)
}
func (us *UserService) Delete(userid string) error {
	return us.UserRepository.Delete(userid)
}
