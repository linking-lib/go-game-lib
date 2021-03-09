package po

import (
	"gorm.io/gorm"
)

type PO interface {
	OnCreate() bool
	NotEmpty() bool
	CacheName() CacheName
}

type AbstractPO struct {
	gorm.Model
}

type CacheName struct {
	Key string
}
