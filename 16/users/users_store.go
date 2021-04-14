package users

import (
	"context"
	"github.com/kirigaikabuto/golangLessons1900/16/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

type UsersStore struct {
	collection *mongo.Collection
}

func NewUsersStore(cf config.MongoConfig) (*UsersStore, error) {
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
	return &UsersStore{collection: collection}, nil
}

func (u *UsersStore) AddUser(user User) error {
	_, err := u.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersStore) List() ([]User, error) {
	users := []User{}
	cursor, err := u.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
