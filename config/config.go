package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App       `yaml:"app"`
		HTTP      `yaml:"http"`
		Shortener `yaml:"shortener"`
		Log       `yaml:"logger"`
		PG        `yaml:"postgres"`
		Repo      `yaml:"repo"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		ServerAddress string `env-required:"true" yaml:"server_address" env:"SERVER_ADDRESS"`
	}

	Shortener struct {
		BaseURL string `env-required:"true" yaml:"base_url" env:"BASE_URL"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// Repo - сохранение в файл
	Repo struct {
		FileStore string `env-required:"true" yaml:"file_store"   env:"FILE_STORAGE_PATH"`
	}

	// PG -.
	PG struct {
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		URL     string `env-required:"true" yaml:"pg_url"  env:"PG_URL"`
	}
)

// NewConfig returns app config.
func NewConfig(address, baseURL, fileStorePath string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	if address != "" {
		cfg.HTTP.ServerAddress = address
	}

	if baseURL != "" {
		cfg.Shortener.BaseURL = baseURL
	}

	if fileStorePath != "" {
		cfg.Repo.FileStore = fileStorePath
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
