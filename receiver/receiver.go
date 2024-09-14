package main

import (
	"fmt"
	"log"
	"rabbitmq/helpers"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	helpers.FailOnError("error connecting to rabbitmq:", err)

	ch, err := conn.Channel()
	helpers.FailOnError("error connecting to rabbitmq:", err)

	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	helpers.FailOnError("error declaring a queue:", err)

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	helpers.FailOnError("error registering a consumer:", err)

	var forever chan struct{}
	go func(){
		for msg := range msgs {
			fmt.Printf(" [x] received: %s\n", msg.Body)
		}
	}()

	log.Println("waiting for messages")
	<-forever
}
