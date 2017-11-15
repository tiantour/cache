package cache

import (
	"fmt"

	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
)

var (
	IP     = "127.0.0.1"
	Port   = "6379"
	p, err = new()
)

// new pool
func new() (*pool.Pool, error) {
	df := func(network, addr string) (*redis.Client, error) {
		return redis.Dial(network, addr)
	}
	account := fmt.Sprintf("%s:%s", IP, Port)
	return pool.NewCustom("tcp", account, 10, df)
}

// operate redis
func operate(cmd string, args ...interface{}) *redis.Resp {
	conn, err := p.Get()
	if err != nil {
		return redis.NewResp(err)
	}
	defer p.Put(conn)
	return conn.Cmd(cmd, args...)
}
