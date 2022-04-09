package main

import (
	"fmt"
	"github.com/whoismarcode/go-chat-room/global"
	"github.com/whoismarcode/go-chat-room/initialize"
)

func main() {
	initialize.LoadConfig("./")

	// test config
	fmt.Println(global.Config.App.RuntimeRootPath)
}
