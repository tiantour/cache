package cache

import "github.com/mediocregopher/radix.v2/redis"

// ZADD 将一个或多个 member 元素及其 score 值加入到有序集 key 当中。
func (z *tZset) ZADD(key string, value ...interface{}) *redis.Resp {
	args := []interface{}{key}
	args = append(args, value...)
	return operate("ZADD", args...)
}

// ZCARD 返回有序集 key 的基数。
func (z *tZset) ZCARD(key string) *redis.Resp {
	return operate("ZCARD", key)
}

// ZCOUNT 返回有序集 key 中， score 值在 min 和 max 之间(默认包括 score 值等于 min 或 max )的成员的数量。
func (z *tZset) ZCOUNT(key, start, stop string) *redis.Resp {
	return operate("ZCOUNT", key, start, stop)
}

// ZINCRBY 为有序集 key 的成员 member 的 score 值加上增量 increment 。
func (z *tZset) ZINCRBY(key string, value ...interface{}) *redis.Resp {
	args := []interface{}{key}
	args = append(args, value...)
	return operate("ZINCRBY", args...)
}

// ZRANGE 返回有序集 key 中，指定区间内的成员。
func (z *tZset) ZRANGE(key, start, stop string, value ...interface{}) *redis.Resp {
	args := []interface{}{key, start, stop}
	args = append(args, value...)
	return operate("ZRANGE", args...)
}

// ZRANGEBYSCORE 返回有序集 key 中，所有 score 值介于 min 和 max 之间(包括等于 min 或 max )的成员。有序集成员按 score 值递增(从小到大)次序排列。
func (z *tZset) ZRANGEBYSCORE(key, start, stop string, value ...interface{}) *redis.Resp {
	args := []interface{}{key, start, stop}
	args = append(args, value...)
	return operate("ZRANGEBYSCORE", args...)
}

// ZRANK 返回有序集 key 中成员 member 的排名。其中有序集成员按 score 值递增(从小到大)顺序排列。
func (z *tZset) ZRANK(key, value string) *redis.Resp {
	return operate("ZRANK", key, value)
}

// ZREM 删除数据
func (z *tZset) ZREM(key string, value ...interface{}) *redis.Resp {
	args := []interface{}{key}
	args = append(args, value...)
	return operate("ZREM", args...)
}

// ZREMRANGEBYRANK 移除有序集 key 中，指定排名(rank)区间内的所有成员。
func (z *tZset) ZREMRANGEBYRANK(key, start, stop string) *redis.Resp {
	return operate("ZREMRANGEBYRANK", key, start, stop)
}

// ZREMRANGEBYSCORE 移除有序集 key 中，所有 score 值介于 min 和 max 之间(包括等于 min 或 max )的成员。
func (z *tZset) ZREMRANGEBYSCORE(key, start, stop string) *redis.Resp {
	return operate("ZREMRANGEBYSCORE", key, start, stop)
}

// ZREVRANGE 返回有序集 key 中，指定区间内的成员。其中成员的位置按 score 值递减(从大到小)来排列。
func (z *tZset) ZREVRANGE(key, start, stop string, value ...interface{}) *redis.Resp {
	args := []interface{}{key, start, stop}
	args = append(args, value...)
	return operate("ZREVRANGE", args...)
}

// ZREVRANGEBYSCORE 返回有序集 key 中， score 值介于 max 和 min 之间(默认包括等于 max 或 min )的所有的成员。有序集成员按 score 值递减(从大到小)的次序排列。
func (z *tZset) ZREVRANGEBYSCORE(key, start, stop string, value ...interface{}) *redis.Resp {
	args := []interface{}{key, start, stop}
	args = append(args, value...)
	return operate("ZREVRANGEBYSCORE", args...)
}

// ZREVRANK 返回有序集 key 中，指定区间内的成员。
func (z *tZset) ZREVRANK(key string, start, stop string, value ...interface{}) *redis.Resp {
	args := []interface{}{key, start, stop}
	args = append(args, value...)
	return operate("ZREVRANK", args...)
}

// ZSCORE 返回有序集 key 中，成员 member 的 score 值。
func (z *tZset) ZSCORE(key, value string) *redis.Resp {
	return operate("ZSCORE", key, value)
}

// ZUNIONSTORE 计算给定的一个或多个有序集的并集，其中给定 key 的数量必须以 numkeys 参数指定，并将该并集(结果集)储存到 destination 。
func (z *tZset) ZUNIONSTORE(dest string, number int, key ...interface{}) *redis.Resp {
	args := []interface{}{dest, number}
	args = append(args, key...)
	return operate("ZUNIONSTORE", args...)
}

// ZINTERSTORE 计算给定的一个或多个有序集的交集，其中给定 key 的数量必须以 numkeys 参数指定，并将该交集(结果集)储存到 destination 。
func (z *tZset) ZINTERSTORE(dest string, number int, key ...interface{}) *redis.Resp {
	args := []interface{}{dest, number}
	args = append(args, key...)
	return operate("ZINTERSTORE", args...)
}

// ZSCAN
func (z *tZset) ZSCAN(key string, cursor int, pattern string, count int) *redis.Resp {
	return operate("ZSCAN", key, cursor, "MATCH", pattern, "COUNT", count)
}

// ZRANGEBYLEX 当有序集合的所有成员都具有相同的分值时， 有序集合的元素会根据成员的字典序（lexicographical ordering）来进行排序， 而这个命令则可以返回给定的有序集合键 key 中， 值介于 min 和 max 之间的成员。
func (z *tZset) ZRANGEBYLEX(key, start, stop string, value ...interface{}) *redis.Resp {
	args := []interface{}{key, start, stop}
	args = append(args, value...)
	return operate("ZRANGEBYLEX", args...)
}

// ZLEXCOUNT 对于一个所有成员的分值都相同的有序集合键 key 来说， 这个命令会返回该集合中， 成员介于 min 和 max 范围内的元素数量。
func (z *tZset) ZLEXCOUNT(key, start, stop string) *redis.Resp {
	return operate("ZLEXCOUNT", key, start, stop)
}

// ZREMRANGEBYLEX 对于一个所有成员的分值都相同的有序集合键 key 来说， 这个命令会移除该集合中， 成员介于 min 和 max 范围内的所有元素。
func (z *tZset) ZREMRANGEBYLEX(key, start, stop string) *redis.Resp {
	return operate("ZREMRANGEBYLEX", key, start, stop)
}
