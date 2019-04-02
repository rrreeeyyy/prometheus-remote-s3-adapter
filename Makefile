VERSION = ${shell cat ./VERSION}
COMMIT = $(shell git describe --always)
PROGNAME = "prometheus-remote-s3-adapter"

build:
	go build -ldflags "-X github.com/rrreeeyyy/prometheus-remote-s3-adapter/cmd.version=${VERSION}" -o ${PROGNAME}
