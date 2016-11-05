package cache

import (
	"fmt"
	"os"

	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
)

// cache
var (
	p, err = new()
	Conn   = conn()
)

// new 链接池
func new() (p *pool.Pool, err error) {
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

// conn
func conn() *redis.Client {
	conn, err := p.Get()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer p.Put(conn)
	return conn
}

// operate
func operate(cmd string, args ...interface{}) *redis.Resp {
	return Conn.Cmd(cmd, args...)
}
