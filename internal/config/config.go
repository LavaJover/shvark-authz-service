package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type AuthzConfig struct {
	Env string 	`yaml:"env"`
	GRPCServer 	`yaml:"grpc_server"`
	AuthzDB 	`yaml:"authz_db"`
	LogConfig 	`yaml:"log_config"`
}

type GRPCServer struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	RetryPolicy	`yaml:"retry_policy"`
}

type RetryPolicy struct {
	MaxAttempts				uint			`yaml:"max_attempts"`
	InitialBackoff			time.Duration	`yaml:"initial_backoff"`
	MaxBackoff				time.Duration	`yaml:"max_backoff"`
	BackoffMultiplier		float32			`yaml:"backoff_multiplier"`
	RetryableStatusCodes	[]string		`yaml:"retryable_status_codes"`
}

type AuthzDB struct {
	Dsn string `yaml:"dsn"`
}

type LogConfig struct {
	LogLevel 	string 	`yaml:"log_level"`
	LogFormat 	string 	`yaml:"log_format"`
	LogOutput 	string 	`yaml:"log_output"`
}

func MustLoad() *AuthzConfig {

	// Processing env config variable and file
	configPath := os.Getenv("AUTHZ_CONFIG_PATH")

	if configPath == ""{
		log.Fatalf("AUTHZ_CONFIG_PATH was not found\n")
	}

	if _, err := os.Stat(configPath); err != nil{
		log.Fatalf("failed to find config file: %v\n", err)
	}

	// YAML to struct object
	var cfg AuthzConfig
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil{
		log.Fatalf("failed to read config file: %v", err)
	}

	return &cfg
}