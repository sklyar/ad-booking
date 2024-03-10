package config

import (
	"fmt"
	"log/slog"
	"strconv"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigyaml"
	"github.com/jackc/pgx/v5/tracelog"
)

type Config struct {
	Logger   Logger   `yaml:"logger" env:"LOGGER"`
	Server   Server   `yaml:"server" env:"SERVER"`
	Database Database `yaml:"database" env:"DATABASE"`
}

type Logger struct {
	Level string `yaml:"level" env:"LEVEL"`
}

func (l Logger) validate() error {
	var level slog.Level
	if err := level.UnmarshalText([]byte(l.Level)); err != nil {
		return fmt.Errorf("level: %w", err)
	}
	return nil
}

type Server struct {
	Host string `yaml:"host" env:"HOST"`
	Port int    `yaml:"port" env:"PORT" required:"true"`
}

func (s Server) Addr() string {
	return s.Host + ":" + strconv.Itoa(s.Port)
}

type Database struct {
	DSN      string `yaml:"dsn" env:"DSN" required:"true"`
	LogLevel string `yaml:"log_level" env:"LOG_LEVEL"`
}

func (d Database) validate() error {
	if d.LogLevel != "" {
		_, err := tracelog.LogLevelFromString(d.LogLevel)
		if err != nil {
			return fmt.Errorf("log_level: %w", err)
		}
	}
	return nil
}

func New(cfgPath string) (*Config, error) {
	var cfg Config
	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		EnvPrefix: "ADB",
		Files:     []string{cfgPath},
		FileDecoders: map[string]aconfig.FileDecoder{
			".yaml": aconfigyaml.New(),
		},
		SkipFlags: true,
	})

	if err := loader.Load(); err != nil {
		return nil, err
	}

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}

	return &cfg, nil
}

func (c *Config) validate() error {
	if err := c.Logger.validate(); err != nil {
		return fmt.Errorf("logger: %w", err)
	}
	if err := c.Database.validate(); err != nil {
		return fmt.Errorf("database: %w", err)
	}
	return nil
}
