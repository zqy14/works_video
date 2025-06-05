package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uint64         `gorm:"column:id;type:bigint(20) UNSIGNED;primaryKey;not null;" json:"id"`
	Mobile    string         `gorm:"column:mobile;type:char(11);comment:手机号;not null;" json:"mobile"`       // 手机号
	Password  string         `gorm:"column:password;type:varchar(32);comment:密码;not null;" json:"password"` // 密码
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
}

func (u *User) TableName() string {
	return "user"
}
