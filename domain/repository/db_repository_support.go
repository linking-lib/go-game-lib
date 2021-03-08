package repository

import (
	"github.com/linking-lib/go-game-lib/common"
	"github.com/linking-lib/go-game-lib/infrastructure/dao"
	"github.com/linking-lib/go-game-lib/infrastructure/po"
	"github.com/linking-lib/go-game-lib/infrastructure/redis"
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

func (db DbRepositorySupport) findFromCache(query interface{}, dest interface{}) bool {
	dbName, key := db.ParseName(query)
	var str = redis.RGet(dbName, key)
	if !strs.IsEmpty(str) {
		common.ParseJson(str, dest)
		return true
	} else {
		return false
	}
}

func (db DbRepositorySupport) find(query interface{}, dest interface{}, dbFind func(dbName string, query interface{}, dest interface{}) int64) bool {
	if db.findFromCache(query, dest) {
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

func (db DbRepositorySupport) save(dbName string, dest interface{}) {
	if dest.(po.PO).OnCreate() {
		db.Dao.InsertOne(dbName, dest)
	} else {
		db.Dao.UpdateOne(dbName, dest)
	}
}

func (db DbRepositorySupport) SaveOne(dest interface{}) {
	dbName, key := db.ParseName(dest)
	redis.RSet(dbName, key, common.ConvertJson(dest))
	db.save(dbName, dest)
}

func (db DbRepositorySupport) SaveList(query interface{}, dest interface{}, values ...interface{}) {
	dbName, key := db.ParseName(query)
	redis.RSet(dbName, key, common.ConvertJson(dest))
	for value := range values {
		db.save(dbName, value)
	}
}
