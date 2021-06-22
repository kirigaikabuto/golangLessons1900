package main

import (
	"encoding/json"
	"fmt"
	"github.com/djumanoff/amqp"
	"github.com/kirigaikabuto/golangLessons1900/29/products"
)

type GetProductByIdCommand struct {
	Id string `json:"id"`
}

func main() {
	//call function getProducts
	//config for rabbitmq
	config := amqp.Config{
		Host:        "localhost",
		VirtualHost: "",
		Port:        5672,
		User:        "",
		Password:    "",
		LogLevel:    5,
	}
	//connect to rabbitmq
	sess := amqp.NewSession(config)
	err := sess.Connect()
	if err != nil {
		panic(err)
		return
	}
	clt, err := sess.Client(amqp.ClientConfig{})
	if err != nil {
		panic(err)
		return
	}
	p1 := &products.Product{}
	response, err := clt.Call("amqp_get_product", amqp.Message{})
	if err != nil {
		panic(err)
		return
	}
	err = json.Unmarshal(response.Body, &p1)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("ID", p1.Id)
	fmt.Println("NAME", p1.Name)
	fmt.Println("PRICE", p1.Price)
	fmt.Println("----EXAMPLE FOR GET PRODUCTS----")
	pr := []products.Product{}
	res, err := clt.Call("amqp_get_products", amqp.Message{})
	if err != nil {
		panic(err)
		return
	}
	err = json.Unmarshal(res.Body, &pr)
	if err != nil {
		panic(err)
		return
	}
	for i, v := range pr {
		fmt.Println(i, v)
	}
	fmt.Println("---EXAMPLE FOR GET PRODUCT BY ID")
	req := &GetProductByIdCommand{Id: "1"}
	reqJson, err := json.Marshal(req)
	if err != nil {
		panic(err)
		return
	}
	product3 := &products.Product{}
	respForProductByID, err := clt.Call("amqp_get_product_by_id", amqp.Message{Body: reqJson})
	if err != nil {
		panic(err)
		return
	}
	err = json.Unmarshal(respForProductByID.Body, &product3)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(product3)
}
