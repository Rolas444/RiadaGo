package core

import (
	// "go/riada/config"
	"time"
)

type Model struct {
	ID        uint      `json:"id" gorm:"primarykey;auto_increment"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

// func (m *Model) Db() *gorm.DB {
// 	return config.Db()
// }
