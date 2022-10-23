package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"sync"
	"time"
)

type PriceTracker struct {
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
			prices, err := prices()
			if err != nil {
				log.Printf("Failed to query prices:%v\n", err)
			}
			p.PricesCh <- prices
		}
	}
}

func prices() (Prices, error) {
	res, err := http.Get("https://api.binance.com/api/v3/ticker/price")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	priceBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var prices Prices
	err = json.Unmarshal(priceBytes, &prices)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	for i, price := range prices {
		priceFloat, _ := new(big.Float).SetString(price.Price)
		prices[i].When = now
		prices[i].PriceFloat = priceFloat
	}
	return prices, nil
}
