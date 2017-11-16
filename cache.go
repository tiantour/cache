package cache

import (
	"fmt"

	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
)

var (
	po  *pool.Pool
	err error
)

// New new cache
func New(ip, port string) {
	if ip == "" {
		ip = "127.0.0.1"
	}
	if port == "" {
		port = ":6379"
	}
	df := func(network, addr string) (*redis.Client, error) {
		return redis.Dial(network, addr)
	}
	po, err = pool.NewCustom("tcp",
		fmt.Sprintf("%s%s", ip, port),
		10,
		df,
	)
}

// operate redis
func operate(cmd string, args ...interface{}) *redis.Resp {
	conn, err := po.Get()
	if err != nil {
		return redis.NewResp(err)
	}
	defer po.Put(conn)
	return conn.Cmd(cmd, args...)
}
