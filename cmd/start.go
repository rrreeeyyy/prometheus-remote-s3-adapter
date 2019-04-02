package cmd

import (
	"github.com/rrreeeyyy/prometheus-remote-s3-adapter/config"
	"github.com/rrreeeyyy/prometheus-remote-s3-adapter/server"

	"github.com/prometheus/common/promlog"
	"github.com/urfave/cli"
)

// StartCommand ...
func StartCommand() cli.Command {
	return cli.Command{
		Name:  "start",
		Usage: "Start prometheus-remote-s3-adapter",
		Action: func(context *cli.Context) error {
			logger := promlog.New(
				&promlog.Config{
					Level:  context.GlobalGeneric("log.level").(*promlog.AllowedLevel),
					Format: context.GlobalGeneric("log.format").(*promlog.AllowedFormat),
				},
			)
			cfg, _ := config.New(context)

			s := server.New(logger, cfg)

			s.Run()
			return nil
		},
	}
}
