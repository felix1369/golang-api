package entity

import (
	"context"

	"github.com/jinzhu/gorm"
)

type Salary struct {
	gorm.Model
	Salary   uint64 `json:"salary"`
	Overtime uint   `json:"overtime"`
	Role     Role
}

// SalaryInfrastructure represent the salary's infrastructure contract
type SalaryInfrastructure interface {
	Fetch(ctx context.Context, query string, args ...interface{}) (res []Salary, err error)
	GetByID(ctx context.Context, id int64) (Salary, error)
	GetByRole(ctx context.Context, title string) (Salary, error)
	Update(ctx context.Context, ar *Salary) error
	Store(ctx context.Context, a *Salary) error
	Delete(ctx context.Context, id int64) error
}

// SalaryApplication represent the salary's application contract
type SalaryApplication interface {
	Fetch(ctx context.Context, query string, args ...interface{}) ([]Salary, string, error)
	GetByID(ctx context.Context, id int64) (Salary, error)
	GetByRole(ctx context.Context, title string) (Salary, error)
	Update(ctx context.Context, ar *Salary) error
	Store(context.Context, *Salary) error
	Delete(ctx context.Context, id int64) error
}
