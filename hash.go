package cache

import "github.com/mediocregopher/radix.v2/redis"

// Hash hash
type Hash struct{}

// NewHash New hash
func NewHash() *Hash {
	return &Hash{}
}

// HDEL 删除哈希表 key 中的一个或多个指定域
func (h Hash) HDEL(key ...interface{}) *redis.Resp {
	return operate("HDEL", key...)
}

// HEXISTS 查看哈希表 key 中，给定域 field 是否存在。
func (h Hash) HEXISTS(key, field string) *redis.Resp {
	return operate("HEXISTS", key, field)
}

// HGET 返回哈希表 key 中给定域 field 的值。
func (h Hash) HGET(key, field string) *redis.Resp {
	return operate("HGET", key, field)
}

// HGETALL 返回哈希表 key 中给定域 field 的值
func (h Hash) HGETALL(key string) *redis.Resp {
	return operate("HGETALL", key)
}

// HINCRBY 为哈希表 key 中的域 field 的值加上增量 increment 。
func (h Hash) HINCRBY(key, field string, value int) *redis.Resp {
	return operate("HINCRBY", key, field, value)
}

// HINCRBYFLOAT 为哈希表 key 中的域 field 加上浮点数增量 increment
func (h Hash) HINCRBYFLOAT(key, field string, value float64) *redis.Resp {
	return operate("HINCRBYFLOAT", key, field, value)
}

// HKEYS 返回哈希表 key 中的所有域。
func (h Hash) HKEYS(key string) *redis.Resp {
	return operate("HKEYS", key)
}

// HLEN 返回哈希表 key 中域的数量。
func (h Hash) HLEN(key string) *redis.Resp {
	return operate("HLEN", key)
}

// HMGET 返回哈希表 key 中，一个或多个给定域的值。
func (h Hash) HMGET(key string, value ...interface{}) *redis.Resp {
	args := []interface{}{key}
	args = append(args, value...)
	return operate("HMGET", args...)
}

// HMSET 同时将多个 field-value (域-值)对设置到哈希表 key 中。
func (h Hash) HMSET(key string, value ...interface{}) *redis.Resp {
	args := []interface{}{key}
	args = append(args, value...)
	return operate("HMSET", args...)
}

// HSET 将哈希表 key 中的域 field 的值设为 value 。
func (h Hash) HSET(key, field, value string) *redis.Resp {
	return operate("HSET", key, field, value)
}

// HSETNX 将哈希表 key 中的域 field 的值设置为 value ，当且仅当域 field 不存在。
func (h Hash) HSETNX(key, field, value string) *redis.Resp {
	return operate("HSETNX", key, field, value)
}

// HVALS 返回哈希表 key 中所有域的值。
func (h Hash) HVALS(key string) *redis.Resp {
	return operate("HVALS", key)
}

// HSCAN scan
func (h Hash) HSCAN(key string, cursor int, pattern string, count int) *redis.Resp {
	return operate("HSCAN", key, cursor, "MATCH", pattern, "COUNT", count)
}

// HSTRLEN str len
func (h Hash) HSTRLEN(key, field string) *redis.Resp {
	return operate("HSTRLEN", key, field)
}
