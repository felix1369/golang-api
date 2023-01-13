package usecase

import (
	"context"
	"time"

	"github.com/felix1369/golang-api/model"
	"github.com/felix1369/golang-api/model/entity"
	"github.com/felix1369/golang-api/model/interfaces"
)

type salaryUsecase struct {
	salaryRepo     interfaces.SalaryRepository
	contextTimeout time.Duration
}

// NewSalaryUsecase will create new an salaryUsecase object representation of domain.SalaryUsecase interface
func NewSalaryUsecase(a interfaces.SalaryRepository, ar entity.Salary, timeout time.Duration) interfaces.SalaryUsecase {
	return &salaryUsecase{
		salaryRepo:     a,
		contextTimeout: timeout,
	}
}

func (a *salaryUsecase) Fetch(c context.Context) (res []entity.Salary, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.salaryRepo.Fetch(ctx, "")
	if err != nil {
		return nil, err
	}

	return
}

func (a *salaryUsecase) GetByID(c context.Context, id int64) (res entity.Salary, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.salaryRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (a *salaryUsecase) Update(c context.Context, ar *entity.Salary) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	ar.UpdatedAt = time.Now()
	return a.salaryRepo.Update(ctx, ar)
}

func (a *salaryUsecase) GetByRole(c context.Context, title string) (res entity.Salary, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	res, err = a.salaryRepo.GetByRole(ctx, title)
	if err != nil {
		return
	}
	return
}

func (a *salaryUsecase) Store(c context.Context, m *entity.Salary) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedSalary, _ := a.GetByRole(ctx, m.Role.Name)
	if existedSalary != (entity.Salary{}) {
		return model.ErrConflict
	}

	err = a.salaryRepo.Store(ctx, m)
	return
}

func (a *salaryUsecase) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedSalary, err := a.salaryRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedSalary == (entity.Salary{}) {
		return model.ErrNotFound
	}
	return a.salaryRepo.Delete(ctx, id)
}
