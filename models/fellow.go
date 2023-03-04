package models

import (
	"go/riada/core"
	"time"
)

type Fellow struct {
	Name      string
	LastName  string
	LastName2 string
	Email     string
	Birthday  *time.Time
	TypeDocID uint
	NDoc      string
	SexID     uint
	Address   Address
	core.Model
}
