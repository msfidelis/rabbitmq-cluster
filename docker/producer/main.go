package main

import (
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

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

	for {

		body := "Hello World!"
		err = channel.Publish(
			"",         // exchannelange
			queue.Name, // routing key
			false,      // mandatory
			false,      // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})

		if err != nil {
			panic("Failed to publish a message:" + err.Error())
		} else {
			log.Printf(" [x] Sent %s", body)
		}

		time.Sleep(time.Second)

	}

}
