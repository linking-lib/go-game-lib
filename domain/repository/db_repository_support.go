package repository

import (
	"github.com/linking-lib/go-game-lib/common"
	"github.com/linking-lib/go-game-lib/infrastructure/dao"
	"github.com/linking-lib/go-game-lib/infrastructure/po"
	"github.com/linking-lib/go-game-lib/infrastructure/redis"
	"github.com/linking-lib/go-game-lib/linking"
	"github.com/linking-lib/go-game-lib/utils/strs"
)

type DbRepositorySupport struct {
	Dao dao.DbDao
}

func (db DbRepositorySupport) FindCache(cacheName po.CacheName, dest interface{}) bool {
	var str = redis.RGet(cacheName.Key)
	if !strs.IsEmpty(str) {
		common.ParseJson(str, dest)
		return true
	} else {
		return false
	}
}

func (db DbRepositorySupport) RemoveCache(cacheName po.CacheName) {
	redis.RDel(cacheName.Key)
}

func (db DbRepositorySupport) FindOne(query interface{}, dest interface{}) bool {
	cacheName := query.(po.PO).CacheName()
	if db.FindCache(cacheName, dest) {
		return true
	}
	if db.Dao.SelectOne(query, dest) > 0 {
		redis.RSet(cacheName.Key, common.ConvertJson(dest))
		return true
	}
	return false
}

func (db DbRepositorySupport) FindList(query interface{}, dest interface{}) bool {
	cacheName := query.(po.PO).CacheName()
	if db.FindCache(cacheName, dest) {
		return true
	}
	if db.Dao.SelectList(query, dest) > 0 {
		redis.RSet(cacheName.Key, common.ConvertJson(dest))
		return true
	}
	return false
}

func (db DbRepositorySupport) Save(cacheName po.CacheName, values ...interface{}) {
	if linking.GetDbCacheMode() == common.DbCacheModeAll {
		db.RemoveCache(cacheName)
	}
	for _, value := range values {
		dest := value.(po.PO)
		if dest.OnCreate() {
			db.Dao.InsertOne(value)
		} else {
			if linking.GetDbCacheMode() == common.DbCacheModeAll {
				db.Dao.UpdateOne(value)
			} else {
				redis.RLRPush(common.DbDataUpdateQueue, cacheName.Key)
			}
		}
	}
}
