package config

import (
	"fmt"
	"os"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Postgresql struct {
		Host     string
		Port     string
		User     string
		DbName   string
		Password string
		SSLmode  string
	}

	Redis struct {
		Host      string
		Port      string
		Password  string
		SecretKey string
	}

	Mailer struct {
		Host     string
		Port     string
		Email    string
		Password string
	}
}

func LoadConfig(env string) *Config {

	cfg := &Config{}

	switch env {
	case "local":
		cfg = loadConfigonLocal()
	case "container":
		cfg = loadConfigonContainer()
	default:
		fmt.Println("Error Loading Config: Incorrect Input")
		os.Exit(2)
	}

	return cfg
}

func loadConfigonLocal() *Config {
	cfg := &Config{}

	yamlFile, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}

func loadConfigonContainer() *Config {
	cfg := &Config{
		Postgresql: struct {
			Host     string
			Port     string
			User     string
			DbName   string
			Password string
			SSLmode  string
		}{
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_DB"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("SSL_MODE"),
		},

		Redis: struct {
			Host      string
			Port      string
			Password  string
			SecretKey string
		}{
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
			os.Getenv("REDIS_PASSWORD"),
			os.Getenv("SECRET_KEY"),
		},

		Mailer: struct {
			Host     string
			Port     string
			Email    string
			Password string
		}{
			os.Getenv("SMTP_HOST"),
			os.Getenv("SMTP_PORT"),
			os.Getenv("EMAIL"),
			os.Getenv("PASSWORD"),
		},
	}

	return cfg
}
