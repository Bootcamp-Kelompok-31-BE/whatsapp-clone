package models

import "gorm.io/gorm"

const (
	UserChatStatus_Pending   = "pending"
	UserChatStatus_Delivered = "delivered"
	UserChatStatus_Read      = "read"
	UserChatStatus_Failed    = "failed"
)

type UserChat struct {
	gorm.Model

	Content  string `gorm:"not null"`
	UserID   uint   `gorm:"not null"`
	TargetID *uint  `gorm:"not null"`
	GroupID  *uint  `gorm:"not null"`
	Status   string `gorm:"not null"`

	User   User   `gorm:"foreignKey:UserID"`
	Target *User  `gorm:"foreignKey:TargetID"`
	Group  *Group `gorm:"foreignKey:GroupID"`
}
