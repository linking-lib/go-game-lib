package dao

import (
	"github.com/linking-lib/go-game-lib/infrastructure/mysql"
)

type DbDao interface {
	SelectOne(query interface{}, dest interface{}) int64
	SelectList(query interface{}, dest interface{}, destList interface{}) int64
	InsertOne(dest interface{}) int64
	UpdateOne(dest interface{}) int64
}

type DbDaoSupport struct {
}

func (db DbDaoSupport) SelectOne(query interface{}, dest interface{}) int64 {
	return mysql.MFindOne(query, dest)
}

func (db DbDaoSupport) SelectList(query interface{}, dest interface{}, destList interface{}) int64 {
	return mysql.MFindList(query, dest, destList)
}

func (db DbDaoSupport) InsertOne(dest interface{}) int64 {
	return mysql.MAdd(dest)
}

func (db DbDaoSupport) UpdateOne(dest interface{}) int64 {
	return mysql.MSave(dest)
}
