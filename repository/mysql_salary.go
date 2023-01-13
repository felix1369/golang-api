package infrastructure

import (
	"context"

	"github.com/felix1369/golang-api/model/entity"
	"github.com/felix1369/golang-api/model/interfaces"
)

type mysqlSalaryInfrastructure struct {
	DB DBHandler
}

// NewMysqlSalary will create an object that represent the salary.Repository interface
func NewMysqlSalary(db DBHandler) interfaces.SalaryRepository {
	return &mysqlSalaryInfrastructure{db}
}

func (m *mysqlSalaryInfrastructure) Fetch(ctx context.Context, query string, args ...interface{}) (res []entity.Salary, err error) {
	if m.DB.Find(&res, "").Error() != nil {
		return nil, m.DB.Error()
	}

	return
}

func (m *mysqlSalaryInfrastructure) GetByID(ctx context.Context, id int64) (res entity.Salary, err error) {
	if m.DB.Find(&res, "").Where("id = ", id).Error() != nil {
		return res, m.DB.Error()
	}

	return
}

func (m *mysqlSalaryInfrastructure) GetByRole(ctx context.Context, role string) (res entity.Salary, err error) {
	if m.DB.Find(&res, "").Where("role = ", role).Error() != nil {
		return res, m.DB.Error()
	}

	return
}

func (m *mysqlSalaryInfrastructure) Store(ctx context.Context, a *entity.Salary) (err error) {
	trx := m.DB.Begin()
	if m.DB.Create(a).Error() != nil {
		trx.Rollback()
		return m.DB.Error()
	}
	trx.Commit()

	return
}

func (m *mysqlSalaryInfrastructure) Delete(ctx context.Context, id int64) (err error) {
	data := entity.Salary{}
	trx := m.DB.Begin()
	if m.DB.Delete(data).Error() != nil {
		trx.Rollback()
		return m.DB.Error()
	}
	trx.Commit()

	return
}
func (m *mysqlSalaryInfrastructure) Update(ctx context.Context, ar *entity.Salary) (err error) {
	trx := m.DB.Begin()
	if m.DB.Delete(ar).Error() != nil {
		trx.Rollback()
		return m.DB.Error()
	}
	trx.Commit()

	return
}
