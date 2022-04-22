package main

import (
	"github.com/whoismarcode/go-chat-room/initializer"
)

func main() {
	initializer.Config("./")
	initializer.Logging()
	initializer.Mysql()
	initializer.Redis()
	initializer.Router()
}
