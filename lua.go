package cache

import (
	"github.com/mediocregopher/radix.v2/redis"
	"github.com/mediocregopher/radix.v2/util"
)

// Eval
func (l *tLua) Eval(script string, keys int, args ...interface{}) *redis.Resp {
	conn, err := p.Get()
	if err != nil {
		return redis.NewResp(err)
	}
	defer p.Put(conn)
	return util.LuaEval(conn, script, keys, args)
}
