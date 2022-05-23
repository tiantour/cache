package cache

import (
	"context"
	"log"
	"time"

	"github.com/mediocregopher/radix/v4"
)

var client radix.Client

func New(Network, Address string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	client, err = (radix.PoolConfig{}).New(ctx, Network, Address)
	if err != nil {
		log.Fatalf("open cache err: %v", err)
	}
}

func do(args radix.Action) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return client.Do(ctx, args)
}
