package cache

import (
	"fmt"
	"os"

	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
)

var (
	p, err = newPool()
)

// newPool 链接池
func newPool() (p *pool.Pool, err error) {
	df := func(network, addr string) (client *redis.Client, err error) {
		client, err = redis.Dial(network, addr)
		return
	}
	return pool.NewCustom("tcp", "127.0.0.1:6379", 10, df)
}

// redis init type
var (
	Key    = &tKey{}
	String = &tString{}
	Hash   = &tHash{}
	List   = &tList{}
	Set    = &tSet{}
	Zset   = &tZset{}
)

// redis type
type (
	tKey    struct{}
	tString struct{}
	tHash   struct{}
	tList   struct{}
	tSet    struct{}
	tZset   struct{}
)

//Operate 操作
func operate(method string, args ...interface{}) *redis.Resp {
	conn, err := p.Get()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer p.Put(conn)
	return conn.Cmd(method, args...)
}
