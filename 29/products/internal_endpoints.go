package products

import (
	"encoding/json"
	"github.com/djumanoff/amqp"
)

//amqp endpoints

type AmqpEndpoints struct {
}

func NewAmqpEndpoints() *AmqpEndpoints {
	return &AmqpEndpoints{}
}

func (a *AmqpEndpoints) GetProductAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		response := GetProduct()
		data, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}
		return &amqp.Message{Body: data}
	}
}

func (a *AmqpEndpoints) GetProductsAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		response := GetProducts()
		data, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}
		return &amqp.Message{Body: data}
	}
}