package biance

import (
	"context"
	"github.com/ljg-cqu/biance/utils/backoff"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

type client struct {
	l log.Logger
	c *http.Client
}

func New(logger log.Logger) Client {
	return &client{logger, &http.Client{}}
}

func (c *client) Do(req *http.Request) (resp *http.Response, err error) {
	err = backoff.RetryFnExponential10Times(c.l, context.Background(), time.Second, time.Second*10, func() (bool, error) {
		resp, err = c.c.Do(req)
		if err != nil {
			return true, errors.Wrapf(err, "failed to send http request")
		}
		return false, nil
	})
	err = errors.WithStack(err)
	return
}
