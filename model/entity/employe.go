package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Employe struct {
	gorm.Model
	Name      string
	Address   string
	BornPlace string
	BornDate  time.Time
	SalaryID  Salary
}
