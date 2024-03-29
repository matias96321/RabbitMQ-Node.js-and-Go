# Node.js + GoLang + RabbitMQ
> implementing a message broker.

Implementing RabbitMQ as an intermediate layer of communication between two simple applications of different languages.

# Producer

Opening the `producer.js` file in the producer folder. this is where we will post messages on RabbitMQ.

```sh
import { connect } from 'amqplib'  
    
const message = "Hello World"
const exchange = "hello_world_exchange"
const queue = "hello_world_queue"
const routingKey = "hello_world_key"

connect("amqp://admin:123456@127.0.0.1:5672/").then((connection)=>{
    
    if(!connection)
    throw new Error('Failed to connect to RabbitMQ')
    
    connection.createChannel().then((channel)=>{

        if(!channel)
        throw new Error('Failed to open a channel')

        channel.assertExchange(exchange, 'direct', { durable: true })

        channel.assertQueue(queue, { durable: true })

        channel.bindQueue(queue, exchange, routingKey)

        channel.publish(exchange, routingKey, Buffer.from(message)) 
        
        console.log("Hello world message sent.");

        setTimeout( function() { channel.close(); connection.close() }, 500 );
    })
})
``` 

In the example above we establish a connection to RabbitMQ using the `amqplib` library, open a channel, declare the exchange and queue, link the queue to the exchange, publish a message to the exchange with the specified routing key, and then close the channel and the connection after a time delay.


# Consumer
 Opening the `consumer.js` file in the consumer folder. This is where we will consume the message sent by the producer.
 
 ```
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
```
In summary, the code establishes a connection to RabbitMQ, creates a communication channel, sets up a consumer to receive messages from the "hello_world_queue", and processes received messages in a separate goroutine. The program runs continuously, waiting for messages from RabbitMQ.


