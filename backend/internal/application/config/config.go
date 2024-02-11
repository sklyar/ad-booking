package config

import (
	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigyaml"
	"strconv"
)

type Config struct {
	Server   Server   `yaml:"server" env:"SERVER"`
	Database Database `yaml:"database" env:"DATABASE"`
}

type Server struct {
	Host string `yaml:"host" env:"HOST"`
	Port int    `yaml:"port" env:"PORT" required:"true"`
}

func (s Server) Addr() string {
	return s.Host + ":" + strconv.Itoa(s.Port)
}

type Database struct {
	DSN string `yaml:"dsn" env:"DSN" required:"true"`
}

func New(cfgPath string) (*Config, error) {
	var cfg Config
	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		EnvPrefix: "ADB",
		Files:     []string{cfgPath},
		FileDecoders: map[string]aconfig.FileDecoder{
			".yaml": aconfigyaml.New(),
		},
	})

	if err := loader.Load(); err != nil {
		return nil, err
	}

	return &cfg, nil
}
