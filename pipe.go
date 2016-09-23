package cache

import "github.com/mediocregopher/radix.v2/redis"

// DISCARD 取消事务，放弃执行事务块内的所有命令。
func (t *tPipe) DISCARD() *redis.Resp {
	return operate("DISCARD")
}

// EXEC 执行所有事务块内的命令。
func (t *tPipe) EXEC() *redis.Resp {
	return operate("EXEC")
}

// MULTI 标记一个事务块的开始。
func (t *tPipe) MULTI() *redis.Resp {
	return operate("MULTI")
}

// UNWATCH 取消 WATCH 命令对所有 key 的监视。
func (t *tPipe) UNWATCH() *redis.Resp {
	return operate("UNWATCH")
}

// WATCH 监视一个(或多个) key ，如果在事务执行之前这个(或这些) key 被其他命令所改动，那么事务将被打断。
func (t *tPipe) WATCH(key ...interface{}) *redis.Resp {
	return operate("WATCH", key...)
}
