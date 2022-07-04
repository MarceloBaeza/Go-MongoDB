package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Users user data
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `json:"name" validate:"required, min=1"`
	Email     string             `json:"email" validate:"required,email"`
	CreatedAt time.Time          `bson:"created_At" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_At" json:"updated_at,omitempty"`
}

// Users user list
type Users []*User

//UserNew struct to add a new user
type UserNew struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}
