package cache

import (
	"fmt"
	"log"

	"github.com/mediocregopher/radix/v3"
	"github.com/tiantour/conf"
)

// client pool
var client *radix.Pool

func init() {
	c := conf.NewCache().Data
	address := fmt.Sprintf("%s%s",
		c.IP,
		c.Port,
	)

	var err error
	client, err = radix.NewPool("tcp", address, 10)
	if err != nil {
		log.Fatalf("open cache err: %v", err)
	}
	defer client.Close()
}

// operate redis
func operate(result interface{}, cmd, key string, args ...interface{}) error {
	action := radix.FlatCmd(result, cmd, key, args...)
	return client.Do(action)
}

// operate redis
func operateS(result interface{}, cmd string, args ...string) error {
	action := radix.Cmd(result, cmd, args...)
	return client.Do(action)
}
