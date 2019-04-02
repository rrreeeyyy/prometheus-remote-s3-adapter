package main

import (
	"os"

	"github.com/rrreeeyyy/prometheus-remote-s3-adapter/cmd"
)

func main() {
	exit := cmd.Run(os.Args)
	os.Exit(exit)
}
