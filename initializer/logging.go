package initializer

import (
	file2 "github.com/whoismarcode/go-chat-room/pkg/file"
	"github.com/whoismarcode/go-chat-room/pkg/logging"
	"log"
)

func Logging() {
	var err error
	filePath := logging.GetLogFilePath()
	fileName := logging.GetLogFileName()
	logging.File, err = file2.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("Logging() err: %v", err)
	}

	logging.Logger = log.New(logging.File, logging.DefaultPrefix, log.LstdFlags)
}