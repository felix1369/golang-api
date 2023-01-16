package entities

import (
	"github.com/jinzhu/gorm"
)

type Salary struct {
	gorm.Model
	Salary       uint64 `json:"salary"`
	Overtime     uint   `json:"overtime"`
	MonthlyBonus uint64 `json:"monthly_bonus"`
	Role         Role
}
