package cache

import (
	"fmt"

	"github.com/mediocregopher/radix/v3"
	"github.com/tiantour/conf"
)

var (
	client *radix.Pool
	err    error
)

func init() {
	c := conf.NewCache().Data
	if c.IP == "" {
		c.IP = "127.0.0.1"
	}
	if c.Port == "" {
		c.Port = ":6379"
	}
	client, err = radix.NewPool("tcp", fmt.Sprintf("%s%s", c.IP, c.Port), 10)
	if err != nil {
		panic(err)
	}
}

// operate redis
func operate(result interface{}, cmd, key string, args ...interface{}) error {
	return client.Do(radix.FlatCmd(result, cmd, key, args...))
}

// operate redis
func operateS(result interface{}, cmd string, args ...string) error {
	return client.Do(radix.Cmd(result, cmd, args...))
}
