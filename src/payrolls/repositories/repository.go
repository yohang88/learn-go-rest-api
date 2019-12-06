package repositories

import (
	"payrolls/entities"
)

type PayrollRepository interface {
	Find(id int) (entities.Payroll, error)
	FindAll() ([]entities.Payroll, error)
	//Update(user *Employee) error
	//Store(user *Employee) (entity.ID, error)
}