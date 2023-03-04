package models

import "go/riada/core"

type TypeDoc struct {
	Name     string
	Desc     string
	Fellow   Fellow
	Comunity Comunity
	core.Model
}
