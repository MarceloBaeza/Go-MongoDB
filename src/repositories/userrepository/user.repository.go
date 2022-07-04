package userrepository

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/mbaezahuenupil/go-mongodb-test/src/database"
	"github.com/mbaezahuenupil/go-mongodb-test/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.Background()
var instace *UserRepository
var once sync.Once

type UserRepository struct {
	collection *mongo.Collection
}

//NewUserRepository constructor
func NewUserRepository(collectionName string) *UserRepository {
	once.Do(func() {
		instace = &UserRepository{
			collection: database.GetCollection(collectionName),
		}
	})
	return instace
}
func (uR *UserRepository) Create(user models.User) error {

	result, err := uR.collection.InsertOne(ctx, user)
	if err != nil {
		fmt.Println("Result Create", result)
		return err
	}
	return nil
}

func (uR *UserRepository) Read() (models.Users, error) {
	var users models.Users

	filter := bson.D{}
	cur, err := uR.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var user models.User
		err = cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil

}
func (uR *UserRepository) Update(user models.User, userId string) error {

	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": oid}
	update := bson.M{
		"$set": bson.M{
			"name":       user.Name,
			"email":      user.Email,
			"updated_At": time.Now(),
		},
	}

	_, err = uR.collection.UpdateOne(ctx, filter, update)

	return err
}
func (uR *UserRepository) Delete(userId string) error {
	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": oid}
	_, err = uR.collection.DeleteOne(ctx, filter)
	return err
}
