package models

import "go/riada/core"

type Ministry struct {
	Name   string
	Desc   string
	Status int `gorm:"status;default:1"`
	core.Model
}
