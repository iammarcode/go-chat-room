package initializer

import (
	"encoding/json"
	"fmt"
	"github.com/iammarcode/go-chat-room/global"
	"github.com/iammarcode/go-chat-room/models"
	"github.com/iammarcode/go-chat-room/pkg/logging"
	"github.com/streadway/amqp"
)

func InitMq() {
	host := global.Config.Mq.Host
	port := global.Config.Mq.Port
	username := global.Config.Mq.Username
	password := global.Config.Mq.Password

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, host, port))
	if err != nil {
		logging.Error("Failed to connect to RabbitMQ, err:", err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		logging.Error("Failed to open a channel, err:", err)
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"chat_room_queue", // name
		false,             // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		logging.Error("Failed to declare a queue, err:", err)
		return
	}

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		logging.Error("Failed to set QoS, err:", err)
		return
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		logging.Error("Failed to register a consumer, err:", err)
	}

	//forever := make(chan bool)

	go func() {
		for msg := range msgs {
			message := models.Message{}
			err = json.Unmarshal(msg.Body, &message)

			if err != nil {
				logging.Error("Failed to unmarshal, err:", err)
			}

			var response error

			switch message.Type {
			case "create":
				response = models.Create(message.Table, message.Data)
			case "update":
				response = models.Update(message.Table, message.Data)
			}

			res, err := json.Marshal(response)
			if err != nil {
				logging.Error("Failed to marshal, err:", err)
			}

			// return res to client
			err = ch.Publish(
				"",          // exchange
				msg.ReplyTo, // routing key
				false,       // mandatory
				false,       // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: msg.CorrelationId,
					Body:          res,
				})
			if err != nil {
				logging.Error("Failed to publish a message, err:", err)
			}

			msg.Ack(false)
		}
	}()

	logging.Info("Awaiting RPC requests...")
	//<-forever
}
