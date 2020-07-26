package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/streadway/amqp"
)

type FakeRegister struct {
	User string `faker:"username"`
	Name string `faker:"first_name"`
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

		message := FakeRegister{}
		err := faker.FakeData(&message)
		body, err := json.Marshal(message)

		if err != nil {
			panic("Failed to parse faker data: " + err.Error())
		}

		err = channel.Publish(
			"",         // exchannelange
			queue.Name, // routing key
			false,      // mandatory
			false,      // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(string(body)),
			})

		if err != nil {
			panic("Failed to publish a message:" + err.Error())
		} else {
			log.Printf("Send a message: %s", string(body))
		}

		time.Sleep(time.Second)

	}

}
