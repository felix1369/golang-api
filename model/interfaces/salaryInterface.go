package interfaces

import (
	"context"

	"github.com/felix1369/golang-api/model/entity"
)

// SalaryInfrastructure represent the salary's infrastructure contract
type SalaryRepository interface {
	Fetch(ctx context.Context, query string, args ...interface{}) (res []entity.Salary, err error)
	GetByID(ctx context.Context, id int64) (entity.Salary, error)
	GetByRole(ctx context.Context, title string) (entity.Salary, error)
	Update(ctx context.Context, ar *entity.Salary) error
	Store(ctx context.Context, a *entity.Salary) error
	Delete(ctx context.Context, id int64) error
}

// SalaryApplication represent the salary's application contract
type SalaryUsecase interface {
	Fetch(ctx context.Context) ([]entity.Salary, error)
	GetByID(ctx context.Context, id int64) (entity.Salary, error)
	GetByRole(ctx context.Context, title string) (entity.Salary, error)
	Update(ctx context.Context, ar *entity.Salary) error
	Store(context.Context, *entity.Salary) error
	Delete(ctx context.Context, id int64) error
}
