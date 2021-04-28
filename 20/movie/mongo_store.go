package movie

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/rand"
	"strings"
	"time"
)

type movieStore struct {
	collection *mongo.Collection
}

func NewMovieStore(config MongoConfig) (MovieStore, error) {
	clientOptions := options.Client().ApplyURI("mongodb://" + config.Host + ":" + config.Port)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	db := client.Database(config.Database)
	err = db.CreateCollection(context.TODO(), config.CollectionName)
	if err != nil && !strings.Contains(err.Error(), "NamespaceExists") {
		return nil, err
	}
	collection := db.Collection(config.CollectionName)
	return &movieStore{collection: collection}, nil
}

func (m *movieStore) Create(movie *Movie) (*Movie, error) {
	rand.Seed(time.Now().UnixNano())
	movie.Id = rand.Intn(1000)
	_, err := m.collection.InsertOne(context.TODO(), movie)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (m *movieStore) Delete(id int) error {
	filter := bson.D{{"id", id}}
	_, err := m.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (m *movieStore) List() ([]Movie, error) {
	movies := []Movie{}
	cursor, err := m.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &movies)
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (m *movieStore) Get(id int) (*Movie, error) {
	filter := bson.D{{"id", id}}
	movie := &Movie{}
	err := m.collection.FindOne(context.TODO(), filter).Decode(&movie)
	if err != nil {
		return nil, err
	}
	return movie, nil
}
