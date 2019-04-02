package client

import (
	"net/http"

	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/prompb"
)

type Client interface {
	Name() string
	String() string
	Shutdown()
}

type Writer interface {
	Write(samples model.Samples, r *http.Request, dryRun bool) ([]byte, error)
	Client
}

type Reader interface {
	Read(req *prompb.ReadRequest, r *http.Request) (*prompb.ReadResponse, error)
	Client
}
