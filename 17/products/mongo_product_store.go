package products

import (
	"context"
	"github.com/kirigaikabuto/golangLessons1900/17/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

//в данной структуре должна находиться поле через которое мы имее прямой доступ к методам добавление и тд в базу
type mongoProductStore struct {
	collection *mongo.Collection
}

//создать глобальный метод который будет создвать подключение к базе

func NewMongoProductStore(config config.MongoConfig) (ProductStore, error) {
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
	return &mongoProductStore{collection: collection}, nil
}

func (mp *mongoProductStore) CreateProduct(product *Product) (*Product, error) {
	return nil, nil
}

func (mp *mongoProductStore) GetProduct(id int64) (*Product, error) {
	return nil, nil
}
