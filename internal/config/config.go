package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env        string `yaml:"env" env:"ENV" env-default:"local"`
	DB         `yaml:"db"`
	HttpServer `yaml:"http_server"`
}

type DB struct {
	Host     string `yaml:"host" env:"DB_HOST" env-required:"true"`
	Port     string `yaml:"port" env:"DB_PORT" env-required:"true"`
	User     string `yaml:"user" env:"DB_USER" env-required:"true"`
	Password string `yaml:"password" env:"DB_PASSWORD" env-required:"true"`
	Name     string `yaml:"name" env:"DB_NAME" env-required:"true"`
}

type HttpServer struct {
	Address     string        `yaml:"address" env:"SERVER_ADDRESS" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env:"SERVER_TIMEOUT" env-default:"4s"`
	IdleTimeout string        `yaml:"idle_timeout" env:"SERVER_IDLE_TIMEOUT" env-default:"60s"`
}

func MustLoad() *Config {
	cfgPath := os.Getenv("CFG_PATH")
	if cfgPath == "" {
		log.Fatal("CFG_PATH is not set")
	}

	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Fatal("config file is not exists:", err)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		log.Fatalf("failed to read config %s: %s", cfgPath, err)
	}

	return &cfg
}
