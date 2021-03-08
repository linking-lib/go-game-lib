package dao

import (
	"github.com/linking-lib/go-game-lib/infrastructure/mysql"
	"github.com/linking-lib/go-game-lib/infrastructure/po"
)

type DbDao interface {
	DbName(value interface{}) string
	CacheName(value interface{}) string
	SelectOne(dbName string, query interface{}, dest interface{})
	SelectList(dbName string, query interface{}, dest interface{})
	InsertOne(dbName string, dest interface{}) int64
	UpdateOne(dbName string, dest interface{}) int64
}

type DbDaoSupport struct {
}

func (db DbDaoSupport) DbName(value interface{}) string {
	return value.(po.PO).DbName()
}

func (db DbDaoSupport) CacheName(value interface{}) string {
	return value.(po.PO).CacheName(value)
}

func (db DbDaoSupport) SelectOne(dbName string, query interface{}, dest interface{}) {
	mysql.MFindOne(dbName, dest, query)
}

func (db DbDaoSupport) SelectList(dbName string, query interface{}, dest interface{}) {
	mysql.MFind(dbName, dest, query)
}

func (db DbDaoSupport) InsertOne(dbName string, dest interface{}) int64 {
	return mysql.MAdd(dbName, dest)
}

func (db DbDaoSupport) UpdateOne(dbName string, dest interface{}) int64 {
	return mysql.MSave(dbName, dest)
}