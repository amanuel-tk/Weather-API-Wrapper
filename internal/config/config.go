package config

import (
	"errors"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Redis   RedisConfig
	Weather WeatherConfig
}

type RedisConfig struct {
	Address  string
	CacheTTL time.Duration
}

type WeatherConfig struct {
	APIKey string
}

func Load() (Config, error) {

	_ = godotenv.Load()

	cfg := Config{
		Redis: RedisConfig{
			Address:  os.Getenv("REDIS_ADDR"),
			CacheTTL: time.Minute * 10,
		},
		Weather: WeatherConfig{
			APIKey: os.Getenv("WEATHER_API_KEY"),
		},
	}

	if cfg.Redis.Address == "" {
		return Config{}, errors.New("REDIS_ADDR is required")
	}

	if cfg.Weather.APIKey == "" {
		return Config{}, errors.New("WEATHER_API_KEY is required")
	}
	return cfg, nil
}
