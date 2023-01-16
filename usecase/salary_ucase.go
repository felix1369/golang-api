package usecase

import (
	"context"
	"time"

	"github.com/felix1369/golang-api/model"
	"github.com/felix1369/golang-api/model/entities"
	"github.com/felix1369/golang-api/model/interfaces"
)

type salaryUsecase struct {
	salaryRepo     interfaces.SalaryRepository
	contextTimeout time.Duration
}

// NewSalaryUsecase will create new an salaryUsecase object representation of domain.SalaryUsecase interface
func NewSalaryUsecase(a interfaces.SalaryRepository, ar entities.Salary, timeout time.Duration) interfaces.SalaryUsecase {
	return &salaryUsecase{
		salaryRepo:     a,
		contextTimeout: timeout,
	}
}

func (a *salaryUsecase) Fetch(c context.Context) (res []entities.Salary, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.salaryRepo.Fetch(ctx, "")
	if err != nil {
		return nil, err
	}

	return
}

func (a *salaryUsecase) GetByID(c context.Context, id uint) (res entities.Salary, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.salaryRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (a *salaryUsecase) Update(c context.Context, ar *entities.Salary) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	ar.UpdatedAt = time.Now()
	return a.salaryRepo.Update(ctx, ar)
}

func (a *salaryUsecase) GetByRoleId(c context.Context, roleId uint) (res entities.Salary, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	res, err = a.salaryRepo.GetByRoleId(ctx, roleId)
	if err != nil {
		return
	}
	return
}

func (a *salaryUsecase) Store(c context.Context, m *entities.Salary) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedSalary, _ := a.GetByRoleId(ctx, m.Role.ID)
	if existedSalary != (entities.Salary{}) {
		return model.ErrConflict
	}

	err = a.salaryRepo.Store(ctx, m)
	return
}

func (a *salaryUsecase) Delete(c context.Context, id uint) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedSalary, err := a.salaryRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedSalary == (entities.Salary{}) {
		return model.ErrNotFound
	}
	return a.salaryRepo.Delete(ctx, id)
}
