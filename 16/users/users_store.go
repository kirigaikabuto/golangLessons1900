package users

import (
	"context"
	"github.com/kirigaikabuto/golangLessons1900/16/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/rand"
	"strings"
	"time"
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
	rand.Seed(time.Now().UnixNano())
	user.Id = rand.Intn(1000)
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

func (u *UsersStore) GetById(id int) (*User, error) {
	filter := bson.D{{"id", id}}
	user := &User{}
	err := u.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UsersStore) Update(user User) error {
	oldUser, err := u.GetById(user.Id)
	if err != nil {
		return err
	}
	if user.Password == "" {
		user.Password = oldUser.Password
	}
	if user.Username == "" {
		user.Username = oldUser.Username
	}
	updateCmd := bson.D{{"$set", bson.D{
		{"username", user.Username},
		{"password", user.Password},
	}}}
	filter := bson.D{{"id", user.Id}}
	_, err = u.collection.UpdateOne(context.TODO(), filter, updateCmd)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsersStore) Delete(id int) error {
	filter := bson.D{{"id", id}}
	_, err := u.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
