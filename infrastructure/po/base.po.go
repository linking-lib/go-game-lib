package po

import (
	"gorm.io/gorm"
)

type PO interface {
	IsCreate() bool
	CacheName() CacheName
}

type AbstractPO struct {
	gorm.Model
}

func (s AbstractPO) IsCreate() bool {
	return s.ID == 0
}

type CacheName struct {
	Key     string
	HashKey string
}
