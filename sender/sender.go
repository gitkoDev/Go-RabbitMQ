package main

import (
	"context"
	"log"
	"time"

	"rabbitmq/helpers"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	helpers.FailOnError("error connecting to rabbitmq server:", err)
	defer conn.Close()

	ch, err := conn.Channel()
	helpers.FailOnError("error opening message channel:", err)
	defer ch.Close()

	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	helpers.FailOnError("error creating queue:", err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World"
	err = ch.PublishWithContext(ctx, "", q.Name, false, false, amqp091.Publishing{ContentType: "text/plain", Body: []byte(body)})
	helpers.FailOnError("error publishing a message:", err)

	log.Printf(" [x] sent: %s\n", body)
}
