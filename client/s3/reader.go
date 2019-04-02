package s3

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/prometheus/prompb"
)

func (c *Client) handleReadQuery(ctx context.Context, query *prompb.Query) (*prompb.QueryResult, error) {
	res := &prompb.QueryResult{}

	from := int(query.StartTimestampMs / 1000)
	until := int(query.EndTimestampMs / 1000)

	if until < from {
		level.Debug(c.logger).Log("msg", "Skipping query with empty time range")
		return res, nil
	}
}

func (c *Client) Read(req *prompb.ReadRequest, r *http.Request) (*prompb.ReadResponse, error) {
	resp := &prompb.ReadResponse{}
	ctx, cancel := context.WithTimeout(context.Background(), c.readTimeout)

	for _, query := range req.Queries {
		queryResult, err := c.handleReadQuery(ctx, query)
		if err != nil {
			return nil, err
		}
		resp.Results = append(resp.Results, queryResult)
	}
	return resp, nil
}
