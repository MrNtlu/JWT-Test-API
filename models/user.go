package models

import (
	"TestAPI/db"
	"TestAPI/requests"
	"TestAPI/utils"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserModel struct {
	Collection *mongo.Collection
}

func NewUserModel(mongoDB *db.MongoDB) *UserModel {
	return &UserModel{
		Collection: mongoDB.Database.Collection("users"),
	}
}

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	EmailAddress string             `bson:"email_address" json:"email_address"`
	Password     string             `bson:"password" json:"-"`
	CreatedAt    time.Time          `bson:"created_at" json:"-"`
}

func createUserObject(emailAddress, password string) *User {
	return &User{
		EmailAddress: emailAddress,
		Password:     utils.HashPassword(password),
		CreatedAt:    time.Now().UTC(),
	}
}

func (userModel *UserModel) CreateUser(data requests.Register) error {
	user := createUserObject(data.EmailAddress, data.Password)

	if _, err := userModel.Collection.InsertOne(context.TODO(), user); err != nil {
		return fmt.Errorf("Failed to create new user.")
	}

	return nil
}

func (userModel *UserModel) FindUserByID(uid string) (User, error) {
	objectUID, _ := primitive.ObjectIDFromHex(uid)

	result := userModel.Collection.FindOne(context.TODO(), bson.M{
		"_id": objectUID,
	})

	var user User
	if err := result.Decode(&user); err != nil {
		return User{}, fmt.Errorf("Failed to find user by id.")
	}

	return user, nil
}

func (userModel *UserModel) FindUserByEmail(email string) (User, error) {
	result := userModel.Collection.FindOne(context.TODO(), bson.M{
		"email_address": email,
	})

	var user User
	if err := result.Decode(&user); err != nil {
		return User{}, fmt.Errorf("Failed to find user by email.")
	}

	return user, nil
}
