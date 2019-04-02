package config

import (
	"time"

	"github.com/prometheus/common/promlog"
	"github.com/urfave/cli"
)

type Config struct {
	LogLevel promlog.AllowedLevel
	Web      webOptions
	Read     readOptions
}

type webOptions struct {
	ListenAddress string
	TelemetryPath string
}

type readOptions struct {
	Timeout     time.Duration
	IgnoreError bool
}

var DefaultConfig = Config{
	Web: webOptions{
		ListenAddress: "0.0.0.0:9201",
		TelemetryPath: "/metrics",
	},
	Read: readOptions{
		Timeout:     5 * time.Minute,
		IgnoreError: true,
	},
}

func New(context *cli.Context) (*Config, error) {
	cfg := &Config{}
	*cfg = DefaultConfig

	return cfg, nil
}
