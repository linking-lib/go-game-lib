package po

import (
	"gorm.io/gorm"
)

type PO interface {
	OnCreate()
	IsCreate() bool
	CacheName() CacheName
}

type AbstractPO struct {
	gorm.Model
	Create bool `gorm:"-"`
}

func (s AbstractPO) OnCreate() {
	s.Create = true
}

func (s AbstractPO) IsCreate() bool {
	return s.Create
}

type CacheName struct {
	Key     string
	HashKey string
}
