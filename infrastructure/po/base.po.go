package po

import (
	"gorm.io/gorm"
)

type PO interface {
	DbName() string
	CacheName(value interface{}) string
	OnCreate() bool
}

type AbstractPO struct {
	gorm.Model
}
