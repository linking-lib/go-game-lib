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

func (db DbRepositorySupport) ParseName(value interface{}) (string, string) {
	var dbName = db.Dao.DbName(value)
	var key = db.Dao.CacheName(value)
	return dbName, key
}

func (db DbRepositorySupport) FindFromCache(query interface{}, dest interface{}) bool {
	dbName, key := db.ParseName(query)
	var str = redis.RGet(dbName, key)
	if !strs.IsEmpty(str) {
		common.ParseJson(str, dest)
		return true
	} else {
		return false
	}
}

func (db DbRepositorySupport) RemoveCache(value interface{}) {
	dbName, key := db.ParseName(value)
	redis.RDel(dbName, key)
}

func (db DbRepositorySupport) find(query interface{}, dest interface{}, dbFind func(dbName string, query interface{}, dest interface{}) int64) bool {
	if db.FindFromCache(query, dest) {
		return true
	}
	dbName, key := db.ParseName(query)
	if dbFind(dbName, query, dest) > 0 {
		redis.RSet(dbName, key, common.ConvertJson(dest))
		return true
	} else {
		return false
	}
}

func (db DbRepositorySupport) FindOne(query interface{}, dest interface{}) bool {
	return db.find(query, dest, db.Dao.SelectOne)
}

func (db DbRepositorySupport) FindList(query interface{}, dest interface{}) bool {
	return db.find(query, dest, db.Dao.SelectList)
}

func (db DbRepositorySupport) save(dest interface{}) {
	dbName, key := db.ParseName(dest)
	if dest.(po.PO).OnCreate() {
		db.Dao.InsertOne(dbName, dest)
	} else {
		if linking.GetDbCacheMode() == common.DbCacheModeAll {
			db.Dao.UpdateOne(dbName, dest)
		} else {
			redis.RLRPush(dbName, common.DbDataUpdateQueue, key)
		}
	}
}

func (db DbRepositorySupport) SaveOne(dest interface{}) {
	if linking.GetDbCacheMode() == common.DbCacheModeAll {
		db.RemoveCache(dest)
	}
	db.save(dest)
}

func (db DbRepositorySupport) SaveList(query interface{}, dest interface{}, values ...interface{}) {
	if linking.GetDbCacheMode() == common.DbCacheModeAll {
		db.RemoveCache(query)
	}
	for value := range values {
		db.save(value)
	}
}
