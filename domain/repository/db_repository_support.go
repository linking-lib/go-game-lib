package repository

import (
	"github.com/linking-lib/go-game-lib/common"
	"github.com/linking-lib/go-game-lib/infrastructure/mysql"
	"github.com/linking-lib/go-game-lib/infrastructure/po"
	"github.com/linking-lib/go-game-lib/infrastructure/redis"
	"github.com/linking-lib/go-game-lib/utils/strs"
)

type DbRepositorySupport struct {
}

func (db DbRepositorySupport) DbName(value interface{}) string {
	return value.(po.PO).DbName()
}

func (db DbRepositorySupport) CacheName(value interface{}) string {
	return value.(po.PO).CacheName(value)
}

func (db DbRepositorySupport) FindOne(query interface{}, dest interface{}) interface{} {
	var dbName = db.DbName(query)
	var key = db.CacheName(query)
	var str = redis.RGet(dbName, key)
	if strs.IsEmpty(str) {
		mysql.MFindOne(dbName, dest, query)
		if dest != nil {
			redis.RSet(dbName, key, common.ConvertJson(dest))
		}
	} else {
		common.ParseJson(str, dest)
	}
	return dest
}

func (db DbRepositorySupport) Find(query interface{}, dest interface{}) interface{} {
	var dbName = db.DbName(query)
	var key = db.CacheName(query)
	var str = redis.RGet(dbName, key)
	if strs.IsEmpty(str) {
		mysql.MFind(dbName, dest, query)
		if dest != nil {
			redis.RSet(dbName, key, common.ConvertJson(dest))
		}
	} else {
		common.ParseJson(str, dest)
	}
	return dest
}

func (db DbRepositorySupport) Save(dest interface{}) {
	var dbName = db.DbName(dest)
	var key = db.CacheName(dest)
	redis.RSet(dbName, key, common.ConvertJson(dest))
}
