package initialize

import (
	logging "github.com/whoismarcode/go-chat-room/logging"
	file "github.com/whoismarcode/go-chat-room/pkg"
	"log"
)

func Logging() {
	var err error
	filePath := logging.GetLogFilePath()
	fileName := logging.GetLogFileName()
	logging.File, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("Logging() err: %v", err)
	}

	logging.Logger = log.New(logging.File, logging.DefaultPrefix, log.LstdFlags)
}