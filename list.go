package cache

import "github.com/mediocregopher/radix.v2/redis"

// BLPOP 它是 LPOP 命令的阻塞版本
func (l *tList) BLPOP(ttl int, key ...interface{}) *redis.Resp {
	args := key
	args = append(args, ttl)
	return operate("BLPOP", args...)
}

// BRPOP 它是 RPOP 命令的阻塞版本
func (l *tList) BRPOP(ttl int, key ...interface{}) *redis.Resp {
	args := key
	args = append(args, ttl)
	return operate("BLPOP", args...)
}

// BRPOPLPUSH BRPOPLPUSH 是 RPOPLPUSH 的阻塞版本
func (l *tList) BRPOPLPUSH(key, dest string, ttl int) *redis.Resp {
	return operate("BRPOPLPUSH", key, dest, ttl)
}

// LINDEX 返回列表 key 中，下标为 index 的元素。
func (l *tList) LINDEX(key string, index int) *redis.Resp {
	return operate("LINDEX", key, index)
}

// LINSERT 将值 value 插入到列表 key 当中，位于值 pivot 之前或之后。
func (l *tList) LINSERT(key, action, pivot, value string) *redis.Resp {
	return operate("LINSERT", key, action, pivot, value)
}

// LLEN 返回列表 key 的长度。
func (l *tList) LLEN(key string) *redis.Resp {
	return operate("LLEN", key)
}

// LPOP 移除并返回列表 key 的头元素。
func (l *tList) LPOP(key string) *redis.Resp {
	return operate("LPOP", key)
}

// LPUSH 将一个或多个值 value 插入到列表 key 的表头
func (l *tList) LPUSH(key string, value ...interface{}) *redis.Resp {
	args := []interface{}{key}
	args = append(args, value...)
	return operate("LPUSH", args...)
}

// LPUSHX 将值 value 插入到列表 key 的表头，当且仅当 key 存在并且是一个列表。
func (l *tList) LPUSHX(key string, value ...interface{}) *redis.Resp {
	args := []interface{}{key}
	args = append(args, value...)
	return operate("LPUSHX", args...)
}

// LRANGE 返回列表 key 中指定区间内的元素，区间以偏移量 start 和 stop 指定。
func (l *tList) LRANGE(key string, start, stop int) *redis.Resp {
	return operate("LRANGE", key, start, stop)
}

// LREM 根据参数 count 的值，移除列表中与参数 value 相等的元素。
// count > 0 : 从表头开始向表尾搜索，移除与 value 相等的元素，数量为 count 。
// count < 0 : 从表尾开始向表头搜索，移除与 value 相等的元素，数量为 count 的绝对值。
// count = 0 : 移除表中所有与 value 相等的值。
func (l *tList) LREM(key, value string, count int) *redis.Resp {
	return operate("LREM", key, count, value)
}

// LSET 将列表 key 下标为 index 的元素的值设置为 value 。
func (l *tList) LSET(key, value string, index int) *redis.Resp {
	return operate("LSET", key, index, value)
}

// LTRIM 对一个列表进行修剪(trim)，就是说，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除。
func (l *tList) LTRIM(key string, start, stop int) *redis.Resp {
	return operate("LTRIM", key, start, stop)
}

// RPOP 移除并返回列表 key 的尾元素。
func (l *tList) RPOP(key string) *redis.Resp {
	return operate("RPOP", key)
}

// RPOPLPUSH
// 将列表 source 中的最后一个元素(尾元素)弹出，并返回给客户端。
// 将 source 弹出的元素插入到列表 destination ，作为 destination 列表的的头元素。
func (l *tList) RPOPLPUSH(key, dest string) *redis.Resp {
	return operate("RPOPLPUSH", key, dest)
}

// RPUSH 将一个或多个值 value 插入到列表 key 的表尾(最右边)。
func (l *tList) RPUSH(key string, value ...interface{}) *redis.Resp {
	args := []interface{}{}
	args = append(args, value...)
	return operate("RPUSH", args...)
}

// RPUSHX 将值 value 插入到列表 key 的表尾，当且仅当 key 存在并且是一个列表。
func (l *tList) RPUSHX(key string, value ...interface{}) *redis.Resp {
	args := []interface{}{}
	args = append(args, value...)
	return operate("RPUSHX", args...)
}
