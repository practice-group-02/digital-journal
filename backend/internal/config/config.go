package config

import (
	"fmt"
	"os"
	"reflect"
)

const (
	ENV     = "env"
	DEFAULT = "default"
)

type Config struct {
	DBconfigs     *DBConfig
	AppConfig *AppConfig
}

type AppConfig struct {
	Port string `env:"APP_PORT" default:"8080"`
}

type DBConfig struct {
	DBhost       string `env:"DB_HOST"`
	DBuser       string `env:"DB_USER"`
	DBpassword   string `env:"DB_PASSWORD"`
	DBname       string `env:"DB_NAME"`
	DBport       string `env:"DB_PORT"`
}


func NewConfig() *Config {
	var dbConfig DBConfig
	var appConfig AppConfig
	v := reflect.ValueOf(dbConfig)
	// db configure
	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get(ENV)
		defaultTag := v.Type().Field(i).Tag.Get(DEFAULT)

		// Skip if tag is not defined or ignored
		if tag == "" || tag == "-" {
			continue
		}
		EnvVar, Info := loadFromEnv(tag, defaultTag)
		a := reflect.Indirect(reflect.ValueOf(dbConfig))
		if Info != "" {
			fmt.Printf("Missing environment configuration for '%s' Loading default setting!\n", a.Type().Field(i).Name)
		}

		reflect.ValueOf(&dbConfig).Elem().Field(i).SetString(EnvVar)
	}

	// app configure
	v = reflect.ValueOf(appConfig)
	tag := v.Type().Field(0).Tag.Get(ENV)
	defaultTag := v.Type().Field(0).Tag.Get(DEFAULT)

	envVar, info := loadFromEnv(tag, defaultTag)
	if info != "" {
		fmt.Printf("miss env port for server")
	}
	appConfig.Port = envVar

	return &Config{
		DBconfigs:  &dbConfig,
		AppConfig: &appConfig,
	}
}

func loadFromEnv(tag string, defaultTag string) (string, string) {
	envVar := os.Getenv(tag)
	if envVar == "" {
		envVar = defaultTag
		return envVar, "1"
	}
	return envVar, ""
}