package config

import (
	"strings"

	"github.com/fitraditya/hook-web/version"
	"github.com/obrel/go-lib/pkg/log"
	"github.com/spf13/viper"
)

func GetEnv() string {
	return viper.GetString("app.env")
}

func GetBaseURL() string {
	if viper.IsSet("base.url") {
		return viper.GetString("base.url")
	}

	return "http://localhost:4000"
}

func EnableRateLimit() bool {
	return viper.GetBool("limiter.enabled")
}

func GetRateLimit() int {
	return viper.GetInt("limiter.rate")
}

func GetMongoURL() string {
	if viper.IsSet("mongo.url") {
		return viper.GetString("mongo.url")
	}

	return "mongodb://127.0.0.1:27017/hook-web"
}

func GetMongoDatabase() string {
	return viper.GetString("mongo.database")
}

func Init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.For("config", "init").Info("No config file found, get config from env var.")
		} else {
			log.For("config", "init").Fatal(err)
		}
	}

	err = log.AddSentryHook("", "", version.Version)
	if err != nil {
		log.For("config", "init").Info("Error while adding sentry hook: %v", err)
	}
}
