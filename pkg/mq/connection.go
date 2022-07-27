package mq

import (
	"fmt"
	"github.com/iammarcode/go-chat-room/global"
	"github.com/iammarcode/go-chat-room/pkg/logging"
	"github.com/streadway/amqp"
)

func GetConn() (*amqp.Connection, error) {
	host := global.Config.Mq.Host
	port := global.Config.Mq.Port
	username := global.Config.Mq.Username
	password := global.Config.Mq.Password

	address := fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, host, port)

	logging.Info("mq dial address:", address)

	conn, err := amqp.Dial(address)
	if err != nil {
		logging.Error("Failed to connect to RabbitMQ, err:", err)
		return nil, err
	}
	defer conn.Close()

	return conn, nil
}
