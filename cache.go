package cache

import (
	"fmt"

	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
	"github.com/tiantour/conf"
)

var (
	p, err = new()
)

// new pool
func new() (*pool.Pool, error) {
	df := func(network, addr string) (*redis.Client, error) {
		return redis.Dial(network, addr)
	}
	account := fmt.Sprintf("%s:%s", conf.Data.Cache.Host, conf.Data.Cache.Port)
	return pool.NewCustom("tcp", account, 10, df)
}

// init type
var (
	Key    = &tKey{}
	String = &tString{}
	Hash   = &tHash{}
	List   = &tList{}
	Set    = &tSet{}
	Zset   = &tZset{}
	Pipe   = &tPipe{}
	Lua    = &tLua{}
)

// redis type
type (
	tKey    struct{}
	tString struct{}
	tHash   struct{}
	tList   struct{}
	tSet    struct{}
	tZset   struct{}
	tPipe   struct{}
	tLua    struct{}
)

// operate redis
func operate(cmd string, args ...interface{}) *redis.Resp {
	conn, err := p.Get()
	if err != nil {
		return redis.NewResp(err)
	}
	defer p.Put(conn)
	return conn.Cmd(cmd, args...)
}
