package infrastructure

import (
	"context"

	"github.com/felix1369/golang-api/model/entities"
	"github.com/felix1369/golang-api/model/interfaces"
)

type mysqlSalaryInfrastructure struct {
	DB DBHandler
}

// NewMysqlSalary will create an object that represent the salary.Repository interface
func NewMysqlSalary(db DBHandler) interfaces.SalaryRepository {
	return &mysqlSalaryInfrastructure{db}
}

func (m *mysqlSalaryInfrastructure) Fetch(ctx context.Context, query string, args ...interface{}) (res []entities.Salary, err error) {
	if m.DB.Find(&res, "").Error() != nil {
		return nil, m.DB.Error()
	}

	return
}

func (m *mysqlSalaryInfrastructure) GetByID(ctx context.Context, id uint) (res entities.Salary, err error) {
	if m.DB.Find(&res, "").Where("id = ", id).Error() != nil {
		return res, m.DB.Error()
	}

	return
}

func (m *mysqlSalaryInfrastructure) GetByRoleId(ctx context.Context, roleId uint) (res entities.Salary, err error) {
	if m.DB.Find(&res, "").Where("role = ", roleId).Error() != nil {
		return res, m.DB.Error()
	}

	return
}

func (m *mysqlSalaryInfrastructure) Store(ctx context.Context, a *entities.Salary) (err error) {
	trx := m.DB.Begin()
	if m.DB.Create(a).Error() != nil {
		trx.Rollback()
		return m.DB.Error()
	}
	trx.Commit()

	return
}

func (m *mysqlSalaryInfrastructure) Delete(ctx context.Context, id uint) (err error) {
	data := entities.Salary{}
	trx := m.DB.Begin()
	if m.DB.Delete(data).Error() != nil {
		trx.Rollback()
		return m.DB.Error()
	}
	trx.Commit()

	return
}
func (m *mysqlSalaryInfrastructure) Update(ctx context.Context, ar *entities.Salary) (err error) {
	trx := m.DB.Begin()
	if m.DB.Delete(ar).Error() != nil {
		trx.Rollback()
		return m.DB.Error()
	}
	trx.Commit()

	return
}
