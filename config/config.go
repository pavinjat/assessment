package config

import "os"

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("missing required environment variable: " + name)
	}
	return v
}

type Config struct {
	Port  string
	DbURL string
}

func NewConfig() *Config {
	return &Config{Port: getenv("PORT"), DbURL: getenv("DATABASE_URL")}
}
