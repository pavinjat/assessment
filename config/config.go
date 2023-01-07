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
	DBurl string
}

func NewConfig() *Config {
	return &Config{Port: getenv("PORT"), DBurl: getenv("DATABASE_URL")}
}
