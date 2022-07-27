package mq

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/iammarcode/go-chat-room/global"
	"github.com/iammarcode/go-chat-room/models"
	"github.com/iammarcode/go-chat-room/pkg/logging"
	"github.com/streadway/amqp"
	"math/rand"
)

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func Pubulish(message models.Message) error {
	host := global.Config.Mq.Host
	port := global.Config.Mq.Port
	username := global.Config.Mq.Username
	password := global.Config.Mq.Password

	address := fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, host, port)

	conn, err := amqp.Dial(address)
	if err != nil {
		logging.Error("Failed to connect to RabbitMQ, err:", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		logging.Error("Failed to open a channel, err:", err)
		return errors.New("sub err")
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		logging.Error("Failed to open a channel, err:", err)
		return errors.New("sub err")
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		logging.Error("Failed to register a consumer, err:", err)
		return errors.New("sub err")
	}

	corrId := randomString(32)

	body, err := json.Marshal(message)
	if err != nil {
		logging.Error("Failed to marshal, err:", err)
		return errors.New("sub err")
	}

	// send request to consumer
	err = ch.Publish(
		"",                // exchange
		"chat_room_queue", // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          body,
		})
	if err != nil {
		logging.Error("Failed to publish a message, err:", err)
		return errors.New("sub err")
	}

	fmt.Println("3")

	// receive response
	go func() error {
		for msg := range msgs {
			fmt.Println("4")
			if corrId == msg.CorrelationId {
				if msg.Body == nil {
					return nil
				}
				var res error
				err = json.Unmarshal(msg.Body, res)
				if err != nil {
					logging.Error("Failed to unmarshal msg.body, err:", err)
					return errors.New("sub err")
				}
				return err
			}
		}
		return nil
	}()

	return errors.New("sub err")
}
