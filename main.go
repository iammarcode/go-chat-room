package main

import (
	"github.com/whoismarcode/go-chat-room/initialize"
	"github.com/whoismarcode/go-chat-room/logging"
)

func init() {
	initialize.LoadConfig("./")
	initialize.Logging()

}

func main() {
	logging.Info("test ar...")
}
