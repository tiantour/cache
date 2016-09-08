package cache

import "github.com/mediocregopher/radix.v2/redis"

// SADD 将一个或多个 member 元素加入到集合 key 当中，已经存在于集合的 member 元素将被忽略。
func (s *tSet) SADD(key string, value ...interface{}) *redis.Resp {
	args := []interface{}{key}
	args = append(args, value...)
	return operate("SADD", args...)
}

// SCARD 返回集合 key 的基数(集合中元素的数量)。
func (s *tSet) SCARD(key string) *redis.Resp {
	return operate("SCARD", key)
}

// SDIFF 返回一个集合的全部成员，该集合是所有给定集合之间的差集。
func (s *tSet) SDIFF(key ...interface{}) *redis.Resp {
	return operate("SDIFF", key...)
}

// SDIFFSTORE 这个命令的作用和 SDIFF 类似，但它将结果保存到 destination 集合，而不是简单地返回结果集。
func (s *tSet) SDIFFSTORE(dest string, key ...interface{}) *redis.Resp {
	args := []interface{}{}
	args = append(args, key...)
	return operate("SDIFFSTORE", args...)
}

// SINTER 返回一个集合的全部成员，该集合是所有给定集合的交集。
func (s *tSet) SINTER(key ...interface{}) *redis.Resp {
	return operate("SINTER", key...)
}

// SINTERSTORE 这个命令类似于 SINTER 命令，但它将结果保存到 destination 集合，而不是简单地返回结果集。
func (s *tSet) SINTERSTORE(dest string, key ...interface{}) *redis.Resp {
	args := []interface{}{}
	args = append(args, key...)
	return operate("SINTERSTORE", args...)
}

// SISMEMBER 判断 member 元素是否集合 key 的成员。
func (s *tSet) SISMEMBER(key, member string) *redis.Resp {
	return operate("SISMEMBER", key, member)
}

// SMEMBERS 返回集合 key 中的所有成员。
func (s *tSet) SMEMBERS(key string) *redis.Resp {
	return operate("SMEMBERS", key)
}

// SMOVE 将 member 元素从 source 集合移动到 destination 集合。
func (s *tSet) SMOVE(key, dest, member string) *redis.Resp {
	return operate("SMOVE", key, dest, member)
}

// SPOP 移除并返回集合中的一个随机元素。
func (s *tSet) SPOP(key string) *redis.Resp {
	return operate("SPOP", key)
}

// SRANDMEMBER 如果命令执行时，只提供了 key 参数，那么返回集合中的一个随机元素。
// 如果 count 为正数，且小于集合基数，那么命令返回一个包含 count 个元素的数组，数组中的元素各不相同。如果 count 大于等于集合基数，那么返回整个集合。
// 如果 count 为负数，那么命令返回一个数组，数组中的元素可能会重复出现多次，而数组的长度为 count 的绝对值。
func (s *tSet) SRANDMEMBER(key string, count int) *redis.Resp {
	return operate("SRANDMEMBER", key, count)
}

// SREM 移除集合 key 中的一个或多个 member 元素，不存在的 member 元素会被忽略。
func (s *tSet) SREM(key ...interface{}) *redis.Resp {
	return operate("SREM", key...)
}

// SUNION 返回一个集合的全部成员，该集合是所有给定集合的并集。
func (s *tSet) SUNION(key ...interface{}) *redis.Resp {
	return operate("SUNION", key...)
}

// SUNIONSTORE 这个命令类似于 SUNION 命令，但它将结果保存到 destination 集合，而不是简单地返回结果集。
func (s *tSet) SUNIONSTORE(dest string, key ...interface{}) *redis.Resp {
	args := []interface{}{dest}
	args = append(args, key...)
	return operate("SUNIONSTORE", args...)
}

// SSCAN
func (s *tSet) SSCAN(key string, cursor int, pattern string, count int) *redis.Resp {
	return operate("SSCAN", key, cursor, "MATCH", pattern, "COUNT", count)
}
