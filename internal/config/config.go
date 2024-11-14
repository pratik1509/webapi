package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	address string
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HttpServer  `yaml:"http_server" env-reuqired:"true"`
}

func MustLoad() *Config {
	var configPath string

	// reading from env variable
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		// reading from inline variable
		flags := flag.String("config", "", "path to configure file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatalf("fatal error - config path not set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("fatal error %s", err.Error())
	}

	var config Config
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("can not read config %s", err.Error())
	}

	return &config
}
