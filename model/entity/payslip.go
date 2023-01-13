package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Payslip struct {
	gorm.Model
	Date      time.Time
	Employe   Employe
	Tax       uint64
	Deduction uint64
}
