package repositories

import (
	"github.com/yohang88/learn-go-rest-api/src/payrolls/entities"
)

type PayrollRepository interface {
	Find(id int) (entities.Payroll, error)
	FindAll() ([]entities.Payroll, error)
	//Update(user *Employee) error
	//Store(user *Employee) (entity.ID, error)
}