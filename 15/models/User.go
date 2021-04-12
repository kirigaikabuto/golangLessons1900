package models

import (
	"context"
	"github.com/kirigaikabuto/golangLessons1900/15/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

type User struct {
	Username string
	Password string
}

func GetConnection(cf config.MongoConfig) (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI("mongodb://" + cf.Host + ":" + cf.Port)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	db := client.Database(cf.Database)
	err = db.CreateCollection(context.TODO(), cf.CollectionName)
	if err != nil && !strings.Contains(err.Error(), "NamespaceExists") {
		return nil, err
	}
	collection := db.Collection(cf.CollectionName)
	return collection, nil
}

func AddUser(collection *mongo.Collection, user User) error {
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func GetAllUsers(collection *mongo.Collection) ([]User, error) {
	users := []User{}
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
