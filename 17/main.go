package main

import (
	"fmt"
	"github.com/kirigaikabuto/golangLessons1900/17/config"
	"github.com/kirigaikabuto/golangLessons1900/17/products"
	"log"
)

func main() {
	productsStore, err := products.NewMongoProductStore(config.MongoConfig{
		Host:           "",
		Port:           "",
		Database:       "",
		CollectionName: "",
	})
	if err != nil {
		log.Fatal(err)
	}
	product, err := productsStore.CreateProduct(&products.Product{
		Id:    0,
		Name:  "",
		Price: 0,
	})
	fmt.Println(product)
}
