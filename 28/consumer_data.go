package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	//connection to rabbitmq
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
		return
	}
	//connection to the channel
	ch, err := connection.Channel()
	if err != nil {
		log.Fatal(err)
		return
	}
	//create or connect to the existed queue(list)
	q, err := ch.QueueDeclare(
		"lesson28",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	messages, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan bool)
	go func() {
		for mes := range messages {
			fmt.Println("Received message is", string(mes.Body))
		}
	}()
	fmt.Println("Waiting message")
	<-forever
}
