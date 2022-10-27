package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/ljg-cqu/biance/backoff"
	"github.com/ljg-cqu/biance/logger"
	"github.com/ljg-cqu/core/smtp"
	"github.com/pkg/errors"
	mail "github.com/xhit/go-simple-mail/v2"

	"io/ioutil"
	"math/big"
	"net/http"
	"sync"
	"time"
)

type PriceTracker struct {
	Logger   logger.Logger
	Interval time.Duration
	PricesCh chan Prices
	WP       *sync.WaitGroup
}

func (p *PriceTracker) Run(ctx context.Context) {
	defer p.WP.Done()
	t := time.NewTicker(p.Interval)
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			prices, err := p.prices(ctx)
			p.Logger.ErrorOnError(err, "Failed to query prices")
			if err != nil {
				email := mail.NewMSG()
				email.SetFrom("Zealy <ljg_cqu@126.com>").
					AddTo("ljg_cqu@126.com").
					SetSubject("Biance Market Price Track Error")
				email.SetBody(mail.TextPlain, fmt.Sprintf("%+v", err))

				err := backoff.RetryFnExponential10Times(p.Logger, ctx, time.Second, time.Second*10, func() (bool, error) {
					emailCli, err := smtp.NewEmailClient(smtp.NetEase126Mail, &tls.Config{InsecureSkipVerify: true}, "ljg_cqu@126.com", "XROTXFGWZUILANPB")
					if err != nil {
						return true, errors.Wrapf(err, "failed to create email client.")
					}
					err = emailCli.Send(email)
					if err != nil {
						return true, errors.Wrapf(err, "failed to send email")
					}
					return false, nil
				})
				p.Logger.ErrorOnError(err, "Failed to report error")
				continue
			}
			p.PricesCh <- prices
		}
	}
}

func (p *PriceTracker) prices(ctx context.Context) (Prices, error) {
	var prices Prices
	err := backoff.RetryFnExponential10Times(p.Logger, ctx, time.Second, time.Second*10, func() (bool, error) {
		res, err := http.Get("https://api.binance.com/api/v3/ticker/price")
		if err != nil {
			return true, errors.Wrapf(err, "failed to send pirce request")
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
