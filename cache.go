package cache

import (
	"fmt"
	"log"

	"github.com/mediocregopher/radix/v3"
	"github.com/tiantour/conf"
)

// cache pool
var cache *radix.Pool

func init() {
	c := conf.NewCache().Data
	address := fmt.Sprintf("%s%s",
		c.IP,
		c.Port,
	)

	var err error
	cache, err = radix.NewPool("tcp", address, 25)
	if err != nil {
		log.Fatalf("open cache err: %v", err)
	}
}

// operate redis
func operate(result interface{}, cmd, key string, args ...interface{}) error {
	action := radix.FlatCmd(result, cmd, key, args...)
	return cache.Do(action)
}

// operate redis
func operateS(result interface{}, cmd string, args ...string) error {
	action := radix.Cmd(result, cmd, args...)
	return cache.Do(action)
}
