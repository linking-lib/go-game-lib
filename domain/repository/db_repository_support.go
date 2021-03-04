package repository

import (
	"github.com/linking-lib/go-game-lib/common"
	"github.com/linking-lib/go-game-lib/infrastructure/mysql"
	po2 "github.com/linking-lib/go-game-lib/infrastructure/po"
	"github.com/linking-lib/go-game-lib/infrastructure/redis"
	"github.com/linking-lib/go-game-lib/utils/strs"
)

type DbRepositorySupport struct {
}

func (db DbRepositorySupport) dbName(value interface{}) string {
	po := value.(po2.AbstractPO)
	return po.Db
}

func (db DbRepositorySupport) cacheName(value interface{}) string {
	po := value.(po2.PO)
	return po.CacheName()
}

func (db DbRepositorySupport) findOne(query interface{}, dest interface{}) {
	var dbName = db.dbName(query)
	var key = db.cacheName(query)
	var str = redis.RGet(dbName, key)
	if strs.IsEmpty(str) {
		mysql.MFindOne(dbName, dest, query)
		redis.RSet(dbName, key, common.ConvertJson(dest))
	} else {
		common.ParseJson(str, dest)
	}
}

func (db DbRepositorySupport) find(query interface{}, dest interface{}) {
	var dbName = db.dbName(query)
	var key = db.cacheName(query)
	var str = redis.RGet(dbName, key)
	if strs.IsEmpty(str) {
		mysql.MFind(dbName, dest, query)
		redis.RSet(dbName, key, common.ConvertJson(dest))
	} else {
		common.ParseJson(str, dest)
	}
}

func (db DbRepositorySupport) save(dest interface{}) {
	var dbName = db.dbName(dest)
	var key = db.cacheName(dest)
	redis.RSet(dbName, key, common.ConvertJson(dest))
}
