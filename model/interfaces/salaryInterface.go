package interfaces

import (
	"context"

	"github.com/felix1369/golang-api/model/entities"
)

// SalaryInfrastructure represent the salary's infrastructure contract
type SalaryRepository interface {
	Fetch(ctx context.Context, query string, args ...interface{}) (res []entities.Salary, err error)
	GetByID(ctx context.Context, id uint) (entities.Salary, error)
	GetByRoleId(ctx context.Context, roleId uint) (entities.Salary, error)
	Update(ctx context.Context, ar *entities.Salary) error
	Store(ctx context.Context, a *entities.Salary) error
	Delete(ctx context.Context, id uint) error
}

// SalaryApplication represent the salary's application contract
type SalaryUsecase interface {
	Fetch(ctx context.Context) ([]entities.Salary, error)
	GetByID(ctx context.Context, id uint) (entities.Salary, error)
	GetByRoleId(ctx context.Context, roleId uint) (entities.Salary, error)
	Update(ctx context.Context, ar *entities.Salary) error
	Store(context.Context, *entities.Salary) error
	Delete(ctx context.Context, id uint) error
}
