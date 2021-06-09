package main

import (
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
	//create producer(send_message)
	message := "Hello my name is yerassyl"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Fatal(err)
		return
	}
}
