package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string        `mapstructure:"env" json:"env" yaml:"env"`
	StoragePath string        `mapstructure:"storage_path" json:"storage_path" yaml:"storage_path"`
	token_ttl   time.Duration `mapstructure:"token_ttl" json:"token_ttl" yaml:"token_ttl"`
	GRPC        GRPCConfig    `mapstructure:"grpc" json:"grpc" yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int `mapstructure:"port" json:"port" yaml:"port"`
	Timeout int `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
}

func MustLoadConfig() *Config {
	path := FetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist")
	}
	var cfg Config
	if err := cleanenv.Load(path, &cfg); err != nil {
		panic(err)
	}
	return &cfg
}

func FetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "config.yaml", "Path to config file")
	flag.Parse()
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}
