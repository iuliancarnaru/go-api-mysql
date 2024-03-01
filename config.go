package main

import (
	"fmt"
	"os"
)

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
	JWTSecret  string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		Port:       getEnv("PORT", "4000"),
		DBUser:     getEnv("MYSQL_USER", "root"),
		DBPassword: getEnv("MYSQL_ROOT_PASSWORD", "password"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("MYSQL_HOST", "127.0.0.1"), getEnv("MYSQL_PORT", "3306")),
		DBName:     getEnv("MYSQL_DATABASE", "projectmanager"),
		JWTSecret:  getEnv("JWT_SECRET", "d8a88f5b6a6d"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
