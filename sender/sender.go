package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gitkoDev/Go-RabbitMQ/helpers"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	helpers.FailOnError("error connecting to rabbitmq server:", err)
	defer conn.Close()

	ch, err := conn.Channel()
	helpers.FailOnError("error opening message channel:", err)
	defer ch.Close()

	q, err := ch.QueueDeclare("hello", true, false, false, false, nil)
	helpers.FailOnError("error creating queue:", err)

	err = ch.Qos(1, 0, false)
	helpers.FailOnError("failed to set Qos", err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := bodyFrom(os.Args)
	err = ch.PublishWithContext(ctx, "", q.Name, false, false, amqp091.Publishing{DeliveryMode: amqp091.Persistent, ContentType: "text/plain", Body: []byte(body)})
	helpers.FailOnError("error publishing a message:", err)

	log.Printf(" [x] sent: %s\n", body)
}

func bodyFrom(args []string) string {
	var s string
	if len(args) < 2 || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

