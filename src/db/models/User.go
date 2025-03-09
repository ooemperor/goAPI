package models

import (
	"fmt"
	"gorm.io/gorm"
)

type AModel struct {
	gorm.Model
}

type User struct {
	AModel
	Name     string `gorm:"unique";serializer:"json"` // A regular string field
	Username string `gorm:"unique";serializer:"json"`
	Email    string // A pointer to a string, allowing for null values
}

func (u User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("before create")
	return
}
