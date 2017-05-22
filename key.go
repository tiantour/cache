package cache

import "github.com/mediocregopher/radix.v2/redis"

// Key key
type Key struct{}

// NewKey new key
func NewKey() *Key {
	return &Key{}
}

// DEL 删除给定的一个或多个 key
func (k Key) DEL(key ...interface{}) *redis.Resp {
	return operate("DEL", key...)
}

// DUMP 序列化给定 key ，并返回被序列化的值
func (k Key) DUMP(key string) *redis.Resp {
	return operate("DUMP", key)
}

// EXISTS 检查给定 key 是否存在。
func (k Key) EXISTS(key string) *redis.Resp {
	return operate("EXISTS", key)
}

// EXPIRE 为给定 key 设置生存时间
func (k Key) EXPIRE(key string, ttl int) *redis.Resp {
	return operate("EXPIRE", key, ttl)
}

// EXPIREAT 为给定 key 设置生存时间 ，参数为时间戳，
func (k Key) EXPIREAT(key string, ttl int) *redis.Resp {
	return operate("EXPIREAT", key, ttl)
}

// KEYS 查找所有符合给定模式 pattern 的 key 。参数支持正则表达式
func (k Key) KEYS(key string) *redis.Resp {
	return operate("KEYS", key)
}

// MIGRATE 将 key 原子性地从当前实例传送到目标实例的指定数据库上，一旦传送成功， key 保证会出现在目标实例上，而当前实例上的 key 会被删除。
func (k Key) MIGRATE(host, prot, key string, dest, timeout int) *redis.Resp {
	return operate("MIGRATE", host, prot, key, dest, timeout)
}

// MOVE 将当前数据库的 key 移动到给定的数据库 db 当中。
func (k Key) MOVE(key string, dest int) *redis.Resp {
	return operate("MOVE", key, dest)
}

// OBJECT 命令允许从内部察看给定 key 的 Redis 对象 , REFCOUNT 和 IDLETIME 返回数字。ENCODING 返回相应的编码类型。
func (k Key) OBJECT(action, key string) *redis.Resp {
	return operate("OBJECT", action, key)
}

// PERSIST 移除给定 key 的生存时间
func (k Key) PERSIST(key string) *redis.Resp {
	return operate("PERSIST", key)
}

// PEXPIRE 为给定 key 设置生存时间 ，单位毫秒
func (k Key) PEXPIRE(key string) *redis.Resp {
	return operate("PEXPIRE", key)
}

// PEXPIREAT 为给定 key 设置生存时间 ，参数为时间戳，单位毫秒
func (k Key) PEXPIREAT(key string) *redis.Resp {
	return operate("PEXPIREAT", key)
}

// PTTL 返回给定 key 的剩余生存时间，单位毫秒
func (k Key) PTTL(key string) *redis.Resp {
	return operate("PTTL", key)
}

// RANDOMKEY 从当前数据库中随机返回(不删除)一个 key
func (k Key) RANDOMKEY(key string) *redis.Resp {
	return operate("RANDOMKEY", key)
}

// RENAME 将给定 key 改名为 newkey 。
func (k Key) RENAME(key, newKey string) *redis.Resp {
	return operate("RENAME", key, newKey)
}

// RENAMENX 当且仅当 newkey 不存在时，将 key 改名为 newkey 。
func (k Key) RENAMENX(key string) *redis.Resp {
	return operate("RENAMENX", key)
}

// RESTORE 反序列化给定的序列化值，并将它和给定的 key 关联。
func (k Key) RESTORE(key, value string, ttl int) *redis.Resp {
	return operate("RESTORE", key, ttl, value)
}

// SORT 返回或保存给定列表、集合、有序集合 key 中经过排序的元素。ASC/DESC ALPHA 字符串
func (k Key) SORT(key string, action ...interface{}) *redis.Resp {
	args := []interface{}{key}
	args = append(args, action...)
	return operate("SORT", args...)
}

// TTL 以秒为单位，返回给定 key 的剩余生存时间
func (k Key) TTL(key string) *redis.Resp {
	return operate("TTL", key)
}

// TYPE 返回 key 所储存的值的类型。
func (k Key) TYPE(key string) *redis.Resp {
	return operate("TYPE", key)
}

// SCAN 迭代当前数据库中的数据库键。
func (k Key) SCAN(key string, cursor int, pattern string, count int) *redis.Resp {
	return operate("SCAN", key, cursor, "MATCH", pattern, "COUNT", count)
}
