package initializer

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/iammarcode/go-chat-room/global"
)

func InitConfig(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("LoadConfig() Fatal error resources file: %w \n", err))
	}
	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("LoadConfig() unable to decode into struct %w \n", err))
	}
}
