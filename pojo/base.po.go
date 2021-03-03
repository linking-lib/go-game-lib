package pojo

import (
	"gorm.io/gorm"
)

type PO interface {
	CacheName() string
}

type AbstractPO struct {
	gorm.Model
	Db string `gorm:"-"`
}
