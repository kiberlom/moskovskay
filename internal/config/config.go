package config

import (
	"fmt"
	"path"

	"github.com/spf13/viper"
)

type Config struct {
	TGToken     string
	TGChatAdmin int
}

func GetConfig() (*Config, error) {
	v := viper.New()
	v.AddConfigPath(path.Join(".."))
	v.SetConfigType("env")
	v.SetConfigName(".env")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("Ошибка чтения конфига: %v", err)
	}

	c := &Config{}

	c.TGToken = v.GetString("TELEGRAM_TOKEN")
	c.TGChatAdmin = v.GetInt("CHAT_ID_ADMIN")

	return c, nil
}
