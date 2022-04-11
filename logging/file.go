package logging

import (
	"fmt"
	"github.com/whoismarcode/go-chat-room/global"
	"time"
)

func GetLogFilePath() string {
	return fmt.Sprintf("%s%s", global.Config.App.RuntimeRootPath, global.Config.Log.LogSavePath)
}

func GetLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		global.Config.Log.LogSaveName,
		time.Now().Format(global.Config.Log.TimeFormat),
		global.Config.Log.LogFileExt,
	)
}