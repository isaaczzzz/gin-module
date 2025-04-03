package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// Config 定义配置结构体
type Config struct {
	GRPC struct {
		Address string
		Port    string
	}
	HTTP struct {
		Port string
	}
}

// LoadConfig 从 .env 文件加载配置
func LoadConfig() (*Config, error) {
	// 加载 .env 文件
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	cfg := &Config{}

	// 加载 HTTP 配置
	cfg.HTTP.Port = os.Getenv("HTTP_PORT")
	if cfg.HTTP.Port == "" {
		cfg.HTTP.Port = ":8080"
	}

	// 加载 gRPC 配置
	cfg.GRPC.Address = os.Getenv("GRPC_ADDRESS")
	if cfg.GRPC.Address == "" {
		cfg.GRPC.Address = "localhost"
	}
	cfg.GRPC.Port = os.Getenv("GRPC_PORT")
	if cfg.GRPC.Port == "" {
		cfg.GRPC.Port = "50051"
	}

	return cfg, nil
}
