package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

const EnvFileName = ".env" // название env файла

// Config структура для парсинга файла конфигурации
type Config struct {
	Server ServerConfig `yaml:"server"`
}

// ServerConfig структура для конфигурации сервера
type ServerConfig struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	Timeout         time.Duration `yaml:"timeout"`
	AccessTokenTTl  time.Duration `yaml:"accessTokenTTl"`
	RefreshTokenTTl time.Duration `yaml:"refreshTokenTTl"`
}

// NewConfig парсит данные из файла конфигурации и возвращает объект Config
func NewConfig(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var conf Config
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

// LoadEnvVariables загружает переменные из env файла
func LoadEnvVariables() error {
	err := godotenv.Load(EnvFileName)
	if err != nil {
		return err
	}

	return nil
}
