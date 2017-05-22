package cache

import "github.com/mediocregopher/radix.v2/redis"

// String string
type String struct{}

// NewString new string
func NewString() *String {
	return &String{}
}

// APPEND 将 value 追加到 key 原来的值的末尾
func (s String) APPEND(key, value string) *redis.Resp {
	return operate("APPEND", key, value)
}

// BITCOUNT 计算给定字符串中，被设置为 1 的比特位的数量。
func (s String) BITCOUNT(key string) *redis.Resp {
	return operate("BITCOUNT", key)
}

// BITOP 对一个或多个保存二进制位的字符串 key 进行位元操作，并将结果保存到 destkey 上。
func (s String) BITOP(action, dest string, key ...interface{}) *redis.Resp {
	args := []interface{}{action, dest}
	args = append(args, key...)
	return operate("BITOP", args...)
}

// BITFIELD BITFIELD 命令可以将一个 Redis 字符串看作是一个由二进制位组成的数组， 并对这个数组中储存的长度不同的整数进行访问 （被储存的整数无需进行对齐）
func (s String) BITFIELD(key string, action ...interface{}) *redis.Resp {
	args := []interface{}{key}
	args = append(args, action...)
	return operate("BITFIELD", args...)
}

// DECR 将 key 中储存的数字值减一。
func (s String) DECR(key string) *redis.Resp {
	return operate("DECR", key)
}

// DECRBY 将 key 所储存的值减去减量 decrement 。
func (s String) DECRBY(key string, value int) *redis.Resp {
	return operate("DECRBY", key, value)
}

// GET 返回 key 所关联的字符串值。
func (s String) GET(key string) (resp *redis.Resp) {
	return operate("GET", key)
}

// GETBIT 对 key 所储存的字符串值，获取指定偏移量上的位(bit)。
func (s String) GETBIT(key string, offset int) (resp *redis.Resp) {
	return operate("GETBIT", key, offset)
}

// GETRANGE 返回 key 中字符串值的子字符串，字符串的截取范围由 start 和 end 两个偏移量决定(包括 start 和 end 在内)。
func (s String) GETRANGE(key, start, stop string) (resp *redis.Resp) {
	return operate("GETRANGE", key, start, stop)
}

// GETSET 将给定 key 的值设为 value ，并返回 key 的旧值(old value)。
func (s String) GETSET(key, value string) (resp *redis.Resp) {
	return operate("GETSET", key, value)
}

// INCR 自增
func (s String) INCR(key string) *redis.Resp {
	return operate("INCR", key)
}

// INCRBY incrby number
func (s String) INCRBY(key string, value int) *redis.Resp {
	return operate("INCRBY", key, value)
}

// INCRBYFLOAT incr by float
func (s String) INCRBYFLOAT(key string, value float64) *redis.Resp {
	return operate("INCRBYFLOAT", key, value)
}

// MGET 返回所有(一个或多个)给定 key 的值。
func (s String) MGET(key ...interface{}) *redis.Resp {
	return operate("MGET", key...)
}

// MSET 同时设置一个或多个 key-value 对。
func (s String) MSET(key ...interface{}) *redis.Resp {
	return operate("MSET", key...)
}

// MSETNX 同时设置一个或多个 key-value 对，当且仅当所有给定 key 都不存在。
func (s String) MSETNX(key ...interface{}) *redis.Resp {
	return operate("MSETNX", key...)
}

// PSETEX 以毫秒为单位设置 key 的生存时间
func (s String) PSETEX(key, value string, ttl int) *redis.Resp {
	return operate("PSETEX", key, ttl, value)
}

// SET 将字符串值 value 关联到 key
func (s String) SET(key, value string, action ...interface{}) (resp *redis.Resp) {
	args := []interface{}{key, value}
	args = append(args, action...)
	return operate("SET", args...)
}

// SETBIT 对 key 所储存的字符串值，设置或清除指定偏移量上的位(bit)。
func (s String) SETBIT(key string, value, offset int) (resp *redis.Resp) {
	return operate("SETBIT", key, offset, value)
}

// SETEX 将值 value 关联到 key ，并将 key 的生存时间设为 seconds (以秒为单位)。
func (s String) SETEX(key, value string, ttl int) (resp *redis.Resp) {
	return operate("SETEX", key, ttl, value)
}

// SETNX 将 key 的值设为 value ，当且仅当 key 不存在。
func (s String) SETNX(key, value string) (resp *redis.Resp) {
	return operate("SETNX", key, value)
}

// SETRANGE 用 value 参数覆写(overwrite)给定 key 所储存的字符串值，从偏移量 offset 开始。
func (s String) SETRANGE(key, value string, offset int) (resp *redis.Resp) {
	return operate("SETRANGE", key, offset, value)
}

// STRLEN 返回 key 所储存的字符串值的长度。
func (s String) STRLEN(key string) (resp *redis.Resp) {
	return operate("STRLEN", key)
}
