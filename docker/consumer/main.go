package main

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func main() {

	url := os.Getenv("AMQP_URL")
	connection, err := amqp.Dial(url)

	if err != nil {
		panic("Failed to connection to RabbitMQ:" + err.Error())
	}

	defer connection.Close()

	channel, err := connection.Channel()

	if err != nil {
		panic("Failed to open a channel to RabbitMQ:" + err.Error())
	}

	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"demo-pub-sub", // name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)

	if err != nil {
		panic("Failed to declare a queue:" + err.Error())
	}

	msgs, err := channel.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)

	if err != nil {
		panic("Failed to consume a queue: " + err.Error())
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			d.Ack(true)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
