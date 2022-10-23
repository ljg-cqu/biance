package main

import (
	"context"
	"encoding/json"
	"github.com/ljg-cqu/biance/backoff"
	"github.com/ljg-cqu/biance/logger"
	"github.com/pkg/errors"

	"io/ioutil"
	"math/big"
	"net/http"
	"sync"
	"time"
)

type PriceTracker struct {
	Logger    logger.Logger
	Interval  time.Duration
	PricesCh  chan Prices
	WaitGroup *sync.WaitGroup
}

func (p *PriceTracker) Run(ctx context.Context) {
	defer p.WaitGroup.Done()
	t := time.NewTicker(p.Interval)
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			prices, err := p.prices(ctx)
			p.Logger.ErrorOnError(err, "Failed to query prices")
			if err != nil {
				continue
			}
			p.PricesCh <- prices
		}
	}
}

func (p *PriceTracker) prices(ctx context.Context) (Prices, error) {
	var prices Prices
	err := backoff.RetryFnExponential100Times(p.Logger, ctx, time.Second, time.Second*10, func() (bool, error) {
		res, err := http.Get("https://api.binance.com/api/v3/ticker/price")
		if err != nil {
			return true, errors.Wrapf(err, "failed to send request")
		}

		defer res.Body.Close()
		priceBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return false, errors.Wrapf(err, "failed to read response body")
		}

		err = json.Unmarshal(priceBytes, &prices)
		if err != nil {
			return false, errors.Wrapf(err, "failed parse price")
		}

		return false, nil
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	now := time.Now()
	for i, price := range prices {
		priceFloat, _ := new(big.Float).SetString(price.Price)
		prices[i].When = now
		prices[i].PriceFloat = priceFloat
	}
	return prices, nil
}
