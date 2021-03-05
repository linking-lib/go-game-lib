package po

import (
	"gorm.io/gorm"
)

type PO interface {
	DbName() string
	CacheName(value interface{}) string
}

type AbstractPO struct {
	gorm.Model
	Db string `gorm:"-"`
}
