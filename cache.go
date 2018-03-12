package cache

import (
	"fmt"

	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
	"github.com/tiantour/conf"
)

var (
	po  *pool.Pool
	err error
)

func init() {
	c := conf.NewConf().Cache
	if c.IP == "" {
		c.IP = "127.0.0.1"
	}
	if c.Port == "" {
		c.Port = ":6379"
	}
	df := func(network, addr string) (*redis.Client, error) {
		return redis.Dial(network, addr)
	}
	po, err = pool.NewCustom("tcp", fmt.Sprintf("%s%s", c.IP, c.Port),
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
