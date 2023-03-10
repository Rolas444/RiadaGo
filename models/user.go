package models

import "go/riada/core"

type User struct {
	UserName string `gorm:"user_name";unique`
	Password string
	Role     string
	Status   int `gorm:"status;default:1"`
	core.Model
}
