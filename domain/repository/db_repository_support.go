package repository

import (
	"github.com/linking-lib/go-game-lib/constants"
	"github.com/linking-lib/go-game-lib/infrastructure/po"
)

type DbRepositorySupport struct {
	Rep DbRepository
}

type DbRepository interface {
	SelectOne(query interface{}, dest interface{}) int64
	SelectList(query interface{}, dest interface{}, destList interface{}) int64
	InsertOne(dest interface{}) int64
	UpdateOne(dest interface{}) int64
	ParseCache(dest interface{}) (string, string)
	ParseCacheList(destList interface{}) []string
	CacheToList(destList []interface{}) interface{}
	ExpireTime() int
}

func (db DbRepositorySupport) ExpireTime() int {
	return constants.Day3Second
}

func (db DbRepositorySupport) FindOne(query interface{}, dest interface{}) bool {
	cacheName := query.(po.PO).CacheName()
	if FindCache(cacheName, dest) {
		return true
	}
	if db.Rep.SelectOne(query, dest) > 0 {
		SaveCache(cacheName, dest, db.Rep.ExpireTime())
		return true
	}
	return false
}

func (db DbRepositorySupport) SaveOne(dest interface{}) {
	value := dest.(po.PO)
	cacheName := value.(po.PO).CacheName()
	// 1、先保存数据库
	if value.IsCreate() {
		db.Rep.InsertOne(dest)
	} else {
		db.Rep.UpdateOne(dest)
	}
	// 2、保存缓存
	SaveCache(cacheName, dest, db.Rep.ExpireTime())
}

func (db DbRepositorySupport) FindList(query interface{}, dest interface{}, destList interface{}) interface{} {
	cacheName := query.(po.PO).CacheName()
	list, size := FindListCache(cacheName, dest, db.Rep)
	if size > 0 {
		return list
	}
	if db.Rep.SelectList(query, dest, destList) > 0 {
		SaveListAllCache(cacheName.Key, destList, db.Rep.ExpireTime(), db.Rep)
	}
	return destList
}

func (db DbRepositorySupport) SaveList(cacheName po.CacheName, values ...interface{}) {
	// 1、先删除缓存
	RemoveCache(cacheName)
	// 2、再修改数据库
	for _, value := range values {
		dest := value.(po.PO)
		if dest.IsCreate() {
			db.Rep.InsertOne(value)
		} else {
			db.Rep.UpdateOne(value)
		}
	}
}
