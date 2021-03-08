package po

import (
	"gorm.io/gorm"
)

type PO interface {
	DbName() string
	CacheOneName(value interface{}) string
	CacheListName(value interface{}) string
	OnCreate() bool
}

type AbstractPO struct {
	gorm.Model
}
