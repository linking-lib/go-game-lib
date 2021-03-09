package po

import (
	"gorm.io/gorm"
)

type PO interface {
	OnCreate() bool
	NotEmpty() bool
}

type AbstractPO struct {
	gorm.Model
}
