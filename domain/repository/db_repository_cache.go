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

func SaveCache(cacheName po.CacheName, dest interface{}) {
	redis.RSet(cacheName.Key, common.ConvertJson(dest))
}

func RemoveCache(cacheName po.CacheName) {
	redis.RDel(cacheName.Key)
}

func FindListCache(cacheName po.CacheName, dest interface{}) []interface{} {
	var list = make([]interface{}, 0)
	strMap := redis.RHGetAll(cacheName.Key)
	if len(strMap) > 0 {
		for _, str := range strMap {
			if !strs.IsEmpty(str) {
				common.ParseJson(str, dest)
				list = append(list, dest)
			}
		}
	}
	return list
}

func SaveListCache(cacheName po.CacheName, dest interface{}) {
	redis.RHSet(cacheName.Key, cacheName.HashKey, common.ConvertJson(dest))
}

func SaveListAllCache(key string, destList interface{}, dbRepository DbRepository) {
	var args = dbRepository.ParseCacheList(destList)
	redis.RHMSet(key, args...)
}

func RemoveListCache(cacheName po.CacheName) {
	redis.RHDel(cacheName.Key, cacheName.HashKey)
}
