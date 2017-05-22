package cache

import "github.com/mediocregopher/radix.v2/redis"

// Pipe pipe
type Pipe struct{}

// NewPipe new pipe
func NewPipe() *Pipe {
	return &Pipe{}
}

// DISCARD 取消事务，放弃执行事务块内的所有命令。
func (p Pipe) DISCARD() *redis.Resp {
	return operate("DISCARD")
}

// EXEC 执行所有事务块内的命令。
func (p Pipe) EXEC() *redis.Resp {
	return operate("EXEC")
}

// MULTI 标记一个事务块的开始。
func (p Pipe) MULTI() *redis.Resp {
	return operate("MULTI")
}

// UNWATCH 取消 WATCH 命令对所有 key 的监视。
func (p Pipe) UNWATCH() *redis.Resp {
	return operate("UNWATCH")
}

// WATCH 监视一个(或多个) key ，如果在事务执行之前这个(或这些) key 被其他命令所改动，那么事务将被打断。
func (p Pipe) WATCH(key ...interface{}) *redis.Resp {
	return operate("WATCH", key...)
}
