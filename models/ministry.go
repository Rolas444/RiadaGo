package models

import "go/riada/core"

type Ministry struct {
	Name   string
	Desc   string
	Status int
	core.Model
}
