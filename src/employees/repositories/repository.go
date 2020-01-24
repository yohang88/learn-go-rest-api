package repositories

import (
	"github.com/yohang88/learn-go-rest-api/src/employees/entities"
)

type EmployeeRepository interface {
	FindAll() ([]entities.Employee, error)
	Find(id int) (entities.Employee, error)
	Store(input *entities.Employee) (entities.Employee, error)
	Update(id int, input *entities.Employee) (entities.Employee, error)
	Destroy(id int) error
}