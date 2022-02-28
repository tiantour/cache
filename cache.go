package cache

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mediocregopher/radix/v4"
	"github.com/tiantour/conf"
)

var address = fmt.Sprintf("%s%s", conf.NewCache().Data.IP, conf.NewCache().Data.Port)

func do(args radix.Action) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := (radix.PoolConfig{}).New(ctx, "tcp", address)
	if err != nil {
		log.Fatalf("open cache err: %v", err)
	}
	defer client.Close()

	return client.Do(ctx, args)
}
