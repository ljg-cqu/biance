package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type ImOnline struct {
	TickDur time.Duration
	WP      *sync.WaitGroup
}

func (i *ImOnline) Tick(ctx context.Context) {
	defer i.WP.Done()
	t := time.NewTicker(i.TickDur)
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			fmt.Println("...")
		}
	}
}
