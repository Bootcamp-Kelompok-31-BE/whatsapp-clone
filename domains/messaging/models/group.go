package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model

	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	ImageURL    string `gorm:"not null"`

	Users []User     `gorm:"many2many:group_users;"`
	Chats []UserChat `gorm:"foreignKey:GroupID"`
}
