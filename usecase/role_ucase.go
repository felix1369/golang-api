package usecase

import (
	"context"
	"time"

	"github.com/felix1369/golang-api/model"
	"github.com/felix1369/golang-api/model/entity"
	"github.com/felix1369/golang-api/model/interfaces"
)

type roleUsecase struct {
	roleRepo       interfaces.RoleRepository
	contextTimeout time.Duration
}

// NewRoleUsecase will create new an roleUsecase object representation of domain.RoleUsecase interface
func NewRoleUsecase(a interfaces.RoleRepository, ar entity.Role, timeout time.Duration) interfaces.RoleUseCase {
	return &roleUsecase{
		roleRepo:       a,
		contextTimeout: timeout,
	}
}

func (a *roleUsecase) Fetch(c context.Context) (res []entity.Role, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.roleRepo.Fetch(ctx, "")
	if err != nil {
		return nil, err
	}

	return
}

func (a *roleUsecase) GetByID(c context.Context, id int64) (res entity.Role, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.roleRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (a *roleUsecase) Update(c context.Context, ar *entity.Role) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	ar.UpdatedAt = time.Now()
	return a.roleRepo.Update(ctx, ar)
}

func (a *roleUsecase) GetByName(c context.Context, title string) (res entity.Role, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	res, err = a.roleRepo.GetByName(ctx, title)
	if err != nil {
		return
	}
	return
}

func (a *roleUsecase) Store(c context.Context, m *entity.Role) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedRole, _ := a.GetByName(ctx, m.Name)
	if existedRole != (entity.Role{}) {
		return model.ErrConflict
	}

	err = a.roleRepo.Store(ctx, m)
	return
}

func (a *roleUsecase) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedRole, err := a.roleRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedRole == (entity.Role{}) {
		return model.ErrNotFound
	}
	return a.roleRepo.Delete(ctx, id)
}
