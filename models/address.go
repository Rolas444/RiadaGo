package models

import "go/riada/core"

type Address struct {
	Desc      string
	Reference string
	Region    string
	Country   string
	FellowID  uint
	core.Model
}
