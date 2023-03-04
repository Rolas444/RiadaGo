package models

import (
	"go/riada/core"
	"time"
)

type Comunity struct {
	Name          string
	LargeName     string
	ShortName     string
	FundationDate *time.Time
	Denomination  string
	TypeDocID     uint
	NDoc          string
	Address       string
	Status        int
	core.Model
}
