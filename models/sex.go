package models

import (
	"go/riada/core"
)

type Sex struct {
	Name      string
	ShortName string
	core.Model
	Fellow Fellow
}
