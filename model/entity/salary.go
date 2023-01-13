package entity

import (
	"github.com/jinzhu/gorm"
)

type Salary struct {
	gorm.Model
	Salary   uint64 `json:"salary"`
	Overtime uint   `json:"overtime"`
	Role     Role
}
