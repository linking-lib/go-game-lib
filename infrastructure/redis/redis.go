/**
 * @Title  redis操作包
 * @Description redis操作的封装
 * @Author YaoWeiXin
 * @Update 2020/11/20 10:09.
 */
package redis

import (
	"github.com/linking-lib/go-game-lib/utils/strs"
)

func Init(redisConfig RConfig) {
	initRedis(redisConfig)
}

func RDel(db string, key string) {
	//c.radius 即为 Circle 类型对象中的属性
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("del", key)
}

func RExists(db string, key string) bool {
	//c.radius 即为 Circle 类型对象中的属性
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("exists", key)
	return rev.(int64) == 1
}

func RExpire(db string, key string, time int) {
	//c.radius 即为 Circle 类型对象中的属性
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("EXPIRE", key, time)
}

func RGet(db string, key string) string {
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("get", key)
	if rev == nil {
		return ""
	} else {
		return strs.ByteToStr(rev)
	}
}

func RSet(db string, key string, value string) {
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("set", key, value)
}

func RSetEX(db string, key string, value string, expire int) {
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("SETEX", key, expire, value)
}

func RSetNX(db string, key string, value string) bool {
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("SETNX", key, value)
	return rev.(int64) == 1
}

func RIncr(db string, key string) {
	RIncrBy(db, key, 1)
}

func RIncrBy(db string, key string, num int) {
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("incrby", key, num)
}

func RHDel(db string, key string, field string) {
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("HDEL", key, field)
}

func RHExists(db string, key string, field string) bool {
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("HEXISTS", key, field)
	return rev.(int64) == 1
}

func RHGet(db string, key string, field string) string {
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("HGET", key, field)
	if rev == nil {
		return ""
	} else {
		return strs.ByteToStr(rev)
	}
}

func RHSet(db string, key string, field string, value string) {
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("HSET", key, field, value)
}

func RHSetNX(db string, key string, field string, value string) bool {
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("HSETNX", key, field, value)
	return rev.(int64) == 1
}

func RHGetAll(db string, key string) map[string]string {
	value := make(map[string]string)
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("HGETALL", key)
	if rev == nil {
		return value
	} else {
		temp := rev.([]interface{})
		num := len(temp)
		for i := 0; i < num; i += 2 {
			k := strs.ByteToStr(temp[i])
			v := strs.ByteToStr(temp[i+1])
			value[k] = v
		}
		return value
	}
}

func RHIncrBy(db string, key string, field string, num int) {
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("HINCRBY", key, field, num)
}

func RHLen(db string, key string) int64 {
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("HLEN", key)
	if rev == nil {
		return 0
	} else {
		return rev.(int64)
	}
}

func RHMGet(db string, key string, field ...string) []string {
	var args = make([]interface{}, 0)
	args = append(args, key)
	for i := range field {
		args = append(args, field[i])
	}
	value := make([]string, 0)
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("HMGET", args...)
	if rev == nil {
		return nil
	} else {
		temp := rev.([]interface{})
		for i := range temp {
			value = append(value, strs.ByteToStr(temp[i]))
		}
		return value
	}
}

func RHMSet(db string, key string, fieldValue ...string) {
	var args = make([]interface{}, 0)
	args = append(args, key)
	for i := range fieldValue {
		args = append(args, fieldValue[i])
	}
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("HMSET", args...)
}

func RHValues(db string, key string) []string {
	value := make([]string, 0)
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("HVALS", key)
	if rev == nil {
		return nil
	} else {
		temp := rev.([]interface{})
		for i := range temp {
			value = append(value, strs.ByteToStr(temp[i]))
		}
		return value
	}
}

func RLLen(db string, key string) int64 {
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("LLEN", key)
	if rev == nil {
		return 0
	} else {
		return rev.(int64)
	}
}

func RLLPop(db string, key string) string {
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("LPOP", key)
	if rev == nil {
		return ""
	} else {
		return strs.ByteToStr(rev)
	}
}

func RLLPush(db string, key string, value ...string) {
	var args = make([]interface{}, 0)
	args = append(args, key)
	for i := range value {
		args = append(args, value[i])
	}
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("LPUSH", args...)
}

func RLRPop(db string, key string) string {
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("RPOP", key)
	if rev == nil {
		return ""
	} else {
		return strs.ByteToStr(rev)
	}
}

func RLRPush(db string, key string, value ...string) {
	var args = make([]interface{}, 0)
	args = append(args, key)
	for i := range value {
		args = append(args, value[i])
	}
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("RPUSH", args...)
}

func RSAdd(db string, key string, value ...string) {
	var args = make([]interface{}, 0)
	args = append(args, key)
	for i := range value {
		args = append(args, value[i])
	}
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("SADD", args...)
}

func RSCard(db string, key string) int64 {
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("SCARD", key)
	if rev == nil {
		return 0
	} else {
		return rev.(int64)
	}
}

func RSRem(db string, key string, value ...string) {
	var args = make([]interface{}, 0)
	args = append(args, key)
	for i := range value {
		args = append(args, value[i])
	}
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("SREM", args...)
}

func RSMembers(db string, key string) []string {
	var args = make([]string, 0)
	var scan int64
	scan = 0
	conn := getConn(db)
	defer releaseConn(conn)
	for true {
		rev, _ := conn.Do("SSCAN", key, scan)
		if rev == nil {
			return nil
		} else {
			value := rev.([]interface{})
			step := strs.ByteToStr(value[0])
			scan = strs.StrToInt64(step)
			list := value[1].([]interface{})
			for i := range list {
				args = append(args, strs.ByteToStr(list[i]))
			}
		}
		if scan == 0 {
			break
		}
	}
	return args
}

/**
redis.RZAdd("default", "test_set", "test1", 1.1)
*/
func RZAdd(db string, key string, member string, score float64) {
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("ZADD", key, score, member)
}

/**
redis.RZAdds("default", "test_set",1.2,  "test2", 1.3, "test3")
*/
func RZAdds(db string, key string, arg ...interface{}) {
	var args = make([]interface{}, 0)
	args = append(args, key)
	for i := range arg {
		args = append(args, arg[i])
	}
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("ZADD", args...)
}

func RZCard(db string, key string) int64 {
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("ZCARD", key)
	if rev == nil {
		return 0
	} else {
		return rev.(int64)
	}
}

func RZIncrBy(db string, key string, member string, score float64) {
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("ZINCRBY", key, score, member)
}

func RZRange(db string, key string, start int, end int, withScore bool, isRev bool) [][]string {
	value := make([][]string, 0)
	conn := getConn(db)
	defer releaseConn(conn)
	command := "ZRANGE"
	if isRev {
		command = "ZREVRANGE"
	}
	var rev interface{}
	if withScore {
		rev, _ = conn.Do(command, key, start, end, "WITHSCORES")
	} else {
		rev, _ = conn.Do(command, key, start, end)
	}
	if rev == nil {
		return nil
	} else {
		temp := rev.([]interface{})
		num := len(temp)
		if withScore {
			for i := 0; i < num; i += 2 {
				k := strs.ByteToStr(temp[i])
				v := strs.ByteToStr(temp[i+1])
				value = append(value, []string{k, v})
			}
		} else {
			for i := 0; i < num; i++ {
				v := strs.ByteToStr(temp[i])
				value = append(value, []string{v})
			}
		}
		return value
	}
}

func RZRangeByScore(db string, key string, min float64, max float64, withScore bool, isRev bool) interface{} {
	return RZRangeByScoreLimit(db, key, min, max, withScore, isRev, -1, -1)
}

func RZRangeByScoreLimit(db string, key string, min float64, max float64, withScore bool, isRev bool, offset int, count int) [][]string {
	value := make([][]string, 0)
	conn := getConn(db)
	defer releaseConn(conn)
	var rev interface{}
	command := "ZRANGEBYSCORE"
	if isRev {
		temp := min
		min = max
		max = temp
		command = "ZREVRANGEBYSCORE"
	}
	if withScore {
		if offset >= 0 && count > 0 {
			rev, _ = conn.Do(command, key, min, max, "WITHSCORES", "LIMIT", offset, count)
		} else {
			rev, _ = conn.Do(command, key, min, max, "WITHSCORES")
		}
	} else {
		if offset >= 0 && count > 0 {
			rev, _ = conn.Do(command, key, min, max, "LIMIT", offset, count)
		} else {
			rev, _ = conn.Do(command, key, min, max)
		}
	}
	if rev == nil {
		return nil
	} else {
		temp := rev.([]interface{})
		num := len(temp)
		if withScore {
			for i := 0; i < num; i += 2 {
				k := strs.ByteToStr(temp[i])
				v := strs.ByteToStr(temp[i+1])
				value = append(value, []string{k, v})
			}
		} else {
			for i := 0; i < num; i++ {
				v := strs.ByteToStr(temp[i])
				value = append(value, []string{v})
			}
		}
		return value
	}
}

/**
0为开始
*/
func RZRank(db string, key string, member string, isRev bool) int64 {
	conn := getConn(db)
	defer releaseConn(conn)
	command := "ZRANK"
	if isRev {
		command = "ZREVRANK"
	}
	rev, _ := conn.Do(command, key, member)
	if rev == nil {
		return 0
	} else {
		return rev.(int64)
	}
}

func RZRem(db string, key string, member ...string) {
	var args = make([]interface{}, 0)
	args = append(args, key)
	for i := range member {
		args = append(args, member[i])
	}
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("ZREM", args...)
}

func RZScore(db string, key string, member string) float64 {
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("ZSCORE", key, member)
	if rev == nil {
		return 0
	} else {
		return strs.StrToFloat(strs.ByteToStr(rev))
	}
}
