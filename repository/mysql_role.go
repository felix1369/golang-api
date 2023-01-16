package infrastructure

import (
	"context"

	"github.com/felix1369/golang-api/model/entities"
	"github.com/felix1369/golang-api/model/interfaces"
)

type mysqlRoleInfrastructure struct {
	DB DBHandler
}

// NewMysqlRole will create an object that represent the role.Repository interface
func NewMysqlRole(db DBHandler) interfaces.RoleRepository {
	return &mysqlRoleInfrastructure{db}
}

func (m *mysqlRoleInfrastructure) Fetch(ctx context.Context, query string, args ...interface{}) (res []entities.Role, err error) {
	if m.DB.Find(&res, "").Error() != nil {
		return nil, m.DB.Error()
	}

	return
}

func (m *mysqlRoleInfrastructure) GetByID(ctx context.Context, id uint) (res entities.Role, err error) {
	if m.DB.Find(&res, "").Where("id = ", id).Error() != nil {
		return res, m.DB.Error()
	}

	return
}

func (m *mysqlRoleInfrastructure) GetByName(ctx context.Context, name string) (res entities.Role, err error) {
	if m.DB.Find(&res, "").Where("name = ", name).Error() != nil {
		return res, m.DB.Error()
	}

	return
}

func (m *mysqlRoleInfrastructure) Store(ctx context.Context, a *entities.Role) (err error) {
	trx := m.DB.Begin()
	if m.DB.Create(a).Error() != nil {
		trx.Rollback()
		return m.DB.Error()
	}
	trx.Commit()

	return
}

func (m *mysqlRoleInfrastructure) Delete(ctx context.Context, id uint) (err error) {
	data := entities.Role{}
	trx := m.DB.Begin()
	if m.DB.Delete(data).Error() != nil {
		trx.Rollback()
		return m.DB.Error()
	}
	trx.Commit()

	return
}
func (m *mysqlRoleInfrastructure) Update(ctx context.Context, ar *entities.Role) (err error) {
	trx := m.DB.Begin()
	if m.DB.Delete(ar).Error() != nil {
		trx.Rollback()
		return m.DB.Error()
	}
	trx.Commit()

	return
}
