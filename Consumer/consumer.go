package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %s", message, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://admin:123456@127.0.0.1:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open channel")
	defer ch.Close()

	msgs, err := ch.Consume(
		"hello_world_queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to consume message")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)
		}
	}()

	fmt.Println("Connected with RabbitMQ")
	fmt.Println("Waiting for messages... ")
	<-forever
}
