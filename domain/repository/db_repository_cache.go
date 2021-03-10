package repository

import (
	"github.com/linking-lib/go-game-lib/common"
	"github.com/linking-lib/go-game-lib/infrastructure/po"
	"github.com/linking-lib/go-game-lib/infrastructure/redis"
	"github.com/linking-lib/go-game-lib/utils/strs"
)

func FindCache(cacheName po.CacheName, dest interface{}) bool {
	var str = redis.RGet(cacheName.Key)
	if !strs.IsEmpty(str) {
		common.ParseJson(str, dest)
		return true
	} else {
		return false
	}
}

func SaveCache(cacheName po.CacheName, dest interface{}, expire int) {
	redis.RSetEX(cacheName.Key, common.ConvertJson(dest), expire)
}

func RemoveCache(cacheName po.CacheName) {
	redis.RDel(cacheName.Key)
}

func FindListCache(cacheName po.CacheName, dest interface{}, dbRepository DbRepository) (interface{}, int) {
	var list = make([]interface{}, 0)
	strMap := redis.RHGetAll(cacheName.Key)
	num := len(strMap)
	if len(strMap) > 0 {
		for _, str := range strMap {
			if !strs.IsEmpty(str) {
				common.ParseJson(str, dest)
				list = append(list, dest)
			}
		}
	}
	return dbRepository.CacheToList(list), num
}

func SaveListCache(cacheName po.CacheName, dest interface{}, expire int) {
	redis.RHSetEx(cacheName.Key, cacheName.HashKey, common.ConvertJson(dest), expire)
}

func SaveListAllCache(key string, destList interface{}, expire int, dbRepository DbRepository) {
	var args = dbRepository.ParseCacheList(destList)
	redis.RHMSetEx(key, expire, args...)
}

func RemoveListCache(cacheName po.CacheName) {
	redis.RHDel(cacheName.Key, cacheName.HashKey)
}
