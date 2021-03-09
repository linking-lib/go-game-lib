package dao

import (
	"github.com/linking-lib/go-game-lib/infrastructure/mysql"
)

type DbDao interface {
	SelectOne(dbName string, query interface{}, dest interface{}) int64
	SelectList(dbName string, query interface{}, dest interface{}) int64
	InsertOne(dbName string, dest interface{}) int64
	UpdateOne(dbName string, dest interface{}) int64
}

type DbDaoSupport struct {
}

func (db DbDaoSupport) SelectOne(dbName string, query interface{}, dest interface{}) int64 {
	return mysql.MFindOne(dbName, dest, query)
}

func (db DbDaoSupport) SelectList(dbName string, query interface{}, dest interface{}) int64 {
	return mysql.MFindList(dbName, dest, query)
}

func (db DbDaoSupport) InsertOne(dbName string, dest interface{}) int64 {
	return mysql.MAdd(dbName, dest)
}

func (db DbDaoSupport) UpdateOne(dbName string, dest interface{}) int64 {
	return mysql.MSave(dbName, dest)
}
