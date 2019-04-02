package cmd

import (
	"fmt"
	"os"

	"github.com/prometheus/common/promlog"
	promlogFlag "github.com/prometheus/common/promlog/flag"
	"github.com/urfave/cli"
)

var (
	version string
)

const (
	exitCodeSuccess = iota
	exitCodeError
	exitCodeMisuse
)

const (
	appName  = "prometheus-remote-s3-adapter"
	appUsage = "Read only featured s3 remote adapter for Prometheus"
)

const (
	levelDefault  = "debug"
	formatDefault = "json"
)

func Run(args []string) int {
	allowedLevel := &promlog.AllowedLevel{}
	allowedFormat := &promlog.AllowedFormat{}
	allowedLevel.Set(levelDefault)
	allowedFormat.Set(formatDefault)

	app := cli.NewApp()
	app.Name = appName
	app.Usage = appUsage
	app.Version = version
	app.Flags = []cli.Flag{
		cli.GenericFlag{
			Name:  promlogFlag.LevelFlagName,
			Value: allowedLevel,
			Usage: promlogFlag.LevelFlagHelp,
		},
		cli.GenericFlag{
			Name:  promlogFlag.FormatFlagName,
			Value: allowedFormat,
			Usage: promlogFlag.FormatFlagHelp,
		},
	}

	app.Commands = []cli.Command{
		StartCommand(),
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitCodeError)
	}

	return 0
}
