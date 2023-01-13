package entity

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	Name  string `json:"name"`
	Level string `json:"level"`
}
