package initializer

import (
	"github.com/iammarcode/go-chat-room/pkg/file"
	"github.com/iammarcode/go-chat-room/pkg/logging"
	"log"
)

func InitLogging() {
	var err error
	filePath := logging.GetLogFilePath()
	fileName := logging.GetLogFileName()
	logging.File, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("Logging() err: %v", err)
	}

	logging.Logger = log.New(logging.File, logging.DefaultPrefix, log.LstdFlags)
}