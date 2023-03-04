package models

import (
	"go/riada/core"
	"time"
)

type Member struct {
	FellowID       uint
	ComunityID     uint
	ComunityOrigin string
	BaptismDate    *time.Time
	MemberShipDate *time.Time
	core.Model
}
