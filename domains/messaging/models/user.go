package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string `gorm:"unique;not null"`

	Chats  []UserChat `gorm:"foreignKey:UserID"`
	Groups []Group    `gorm:"many2many:group_users;"`
}
