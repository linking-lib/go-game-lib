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

type DbRepositoryKey struct {
	DbName   string
	CacheKey string
}

func (db DbRepositorySupport) FindFromCache(key DbRepositoryKey, dest interface{}) bool {
	var str = redis.RGet(key.DbName, key.CacheKey)
	if !strs.IsEmpty(str) {
		common.ParseJson(str, dest)
		return true
	} else {
		return false
	}
}

func (db DbRepositorySupport) RemoveCache(key DbRepositoryKey) {
	redis.RDel(key.DbName, key.CacheKey)
}

func (db DbRepositorySupport) find(key DbRepositoryKey, query interface{}, dest interface{},
	dbFind func(dbName string, query interface{}, dest interface{}) int64) bool {
	if db.FindFromCache(key, dest) {
		return true
	}
	if dbFind(key.DbName, query, dest) > 0 {
		redis.RSet(key.DbName, key.CacheKey, common.ConvertJson(dest))
		return true
	} else {
		return false
	}
}

func (db DbRepositorySupport) FindOne(key DbRepositoryKey, query interface{}, dest interface{}) bool {
	return db.find(key, query, dest, db.Dao.SelectOne)
}

func (db DbRepositorySupport) FindList(key DbRepositoryKey, query interface{}, dest interface{}) bool {
	return db.find(key, query, dest, db.Dao.SelectList)
}

func (db DbRepositorySupport) Save(key DbRepositoryKey, values ...interface{}) {
	if linking.GetDbCacheMode() == common.DbCacheModeAll {
		db.RemoveCache(key)
	}
	for _, value := range values {
		dest := value.(po.PO)
		if dest.OnCreate() {
			db.Dao.InsertOne(key.DbName, value)
		} else {
			if linking.GetDbCacheMode() == common.DbCacheModeAll {
				db.Dao.UpdateOne(key.DbName, value)
			} else {
				redis.RLRPush(key.DbName, common.DbDataUpdateQueue, key.CacheKey)
			}
		}
	}
}
