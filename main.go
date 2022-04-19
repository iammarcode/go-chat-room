package main

import (
	"github.com/whoismarcode/go-chat-room/initialize"
	"github.com/whoismarcode/go-chat-room/logging"
)

func init() {
	initialize.LoadConfig("./")
	initialize.Logging()
	initialize.Mysql()
	initialize.Redis()

}

func main() {
	logging.Info("test ar...")
}
