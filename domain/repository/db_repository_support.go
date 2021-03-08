package repository

import (
	"github.com/linking-lib/go-game-lib/common"
	"github.com/linking-lib/go-game-lib/infrastructure/dao"
	"github.com/linking-lib/go-game-lib/infrastructure/mysql"
	"github.com/linking-lib/go-game-lib/infrastructure/po"
	"github.com/linking-lib/go-game-lib/infrastructure/redis"
	"github.com/linking-lib/go-game-lib/utils/strs"
	"github.com/linking-lib/go-game-lib/utils/util"
)

type DbRepositorySupport struct {
	Dao dao.DbDao
}

func (db DbRepositorySupport) ParseName(value interface{}) (string, string) {
	var dbName = db.Dao.DbName(value)
	var key = db.Dao.CacheName(value)
	return dbName, key
}

func (db DbRepositorySupport) findFromCache(query interface{}, dest interface{}) {
	dbName, key := db.ParseName(query)
	var str = redis.RGet(dbName, key)
	if !strs.IsEmpty(str) {
		common.ParseJson(str, dest)
	}
}

func (db DbRepositorySupport) find(query interface{}, dest interface{}, dbFind func(dbName string, query interface{}, dest interface{})) {
	db.findFromCache(query, dest)
	if util.IsNil(dest) {
		dbName, key := db.ParseName(query)
		dbFind(dbName, query, dest)
		mysql.MFindOne(dbName, dest, query)
		if util.IsNotNil(dest) {
			redis.RSet(dbName, key, common.ConvertJson(dest))
		}
	}
}

func (db DbRepositorySupport) FindOne(query interface{}, dest interface{}) {
	db.find(query, dest, db.Dao.SelectOne)
}

func (db DbRepositorySupport) FindList(query interface{}, dest interface{}) {
	db.find(query, dest, db.Dao.SelectList)
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
