package s3

import (
	"time"

	"github.com/go-kit/kit/log"
)

type Client struct {
	readTimeout time.Duration

	logger log.Logger
}
