package repository

import (
	"github.com/ganeryao/linking-go-agile/common"
	"github.com/ganeryao/linking-go-agile/infrastructure/mysql"
	"github.com/ganeryao/linking-go-agile/infrastructure/redis"
	"github.com/ganeryao/linking-go-agile/pojo"
	"github.com/ganeryao/linking-go-agile/utils/strs"
)

type DbRepositorySupport struct {
}

func (db DbRepositorySupport) dbName(value interface{}) string {
	po := value.(pojo.AbstractPO)
	return po.Db
}

func (db DbRepositorySupport) cacheName(value interface{}) string {
	po := value.(pojo.PO)
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
