package main

import (
	"github.com/iammarcode/go-chat-room/initializer"
)

func main() {
	initializer.InitConfig("./")
	initializer.InitLogging()
	initializer.InitMysql()
	initializer.InitRedis()
	initializer.InitMq()
	initializer.InitRouter()
}
