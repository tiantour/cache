package cache

import (
	"github.com/mediocregopher/radix.v2/redis"
	"github.com/mediocregopher/radix.v2/util"
)

// Eval
func (l *tLua) Eval(script string, keys int, args ...interface{}) *redis.Resp {
	return util.LuaEval(Conn, script, keys, args)
}
