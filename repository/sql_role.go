package repository

import (
	"context"

	"github.com/felix1369/golang-api/model/entities"
	"github.com/felix1369/golang-api/model/interfaces"
)

type sqlRoleInfrastructure struct {
	DB DBHandler
}

// NewSqlRole will create an object that represent the role.Repository interface
func NewSqlRole(db DBHandler) interfaces.RoleRepository {
	return &sqlRoleInfrastructure{db}
}

func (m *sqlRoleInfrastructure) Fetch(ctx context.Context, query string, args ...interface{}) (res []entities.Role, err error) {
	if m.DB.Find(&res, "").Error() != nil {
		return nil, m.DB.Error()
	}

	return
}

func (m *sqlRoleInfrastructure) GetByID(ctx context.Context, id uint) (res entities.Role, err error) {
	if m.DB.Find(&res, "").Where("id = ", id).Error() != nil {
		return res, m.DB.Error()
	}

	return
}

func (m *sqlRoleInfrastructure) GetByName(ctx context.Context, name string) (res entities.Role, err error) {
	if m.DB.Find(&res, "").Where("name = ", name).Error() != nil {
		return res, m.DB.Error()
	}

	return
}

func (m *sqlRoleInfrastructure) Store(ctx context.Context, a *entities.Role) (err error) {
	trx := m.DB.Begin()
	if m.DB.Create(a).Error() != nil {
		trx.Rollback()
		return m.DB.Error()
	}
	trx.Commit()

	return
}

func (m *sqlRoleInfrastructure) Delete(ctx context.Context, id uint) (err error) {
	data := entities.Role{}
	trx := m.DB.Begin()
	if m.DB.Delete(data).Error() != nil {
		trx.Rollback()
		return m.DB.Error()
	}
	trx.Commit()

	return
}
func (m *sqlRoleInfrastructure) Update(ctx context.Context, ar *entities.Role) (err error) {
	trx := m.DB.Begin()
	if m.DB.Delete(ar).Error() != nil {
		trx.Rollback()
		return m.DB.Error()
	}
	trx.Commit()

	return
}
