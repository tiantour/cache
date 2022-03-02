package cache

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mediocregopher/radix/v4"
	"github.com/tiantour/conf"
)

var (
	client  radix.Client
	address = fmt.Sprintf("%s%s", conf.NewCache().Data.IP, conf.NewCache().Data.Port)
)

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	client, err = (radix.PoolConfig{}).New(ctx, "tcp", address)
	if err != nil {
		defer client.Close()
		log.Fatalf("open cache err: %v", err)
	}
}

func do(args radix.Action) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return client.Do(ctx, args)
}
