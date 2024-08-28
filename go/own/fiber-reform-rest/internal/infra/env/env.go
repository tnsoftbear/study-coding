package env

import (
	"github.com/spf13/viper"
)

func Setup() {
	viper.AutomaticEnv()
}

func GetEnv(key, def string) string {
	val := viper.GetString(key)
	if val == "" {
		return def
	}
	return val
}
