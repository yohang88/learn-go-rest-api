package repositories

import (
	"employees/entities"
)

type Repository interface {
	Find(id int) (*entities.Employee, error)
	FindAll() ([]*entities.Employee, error)
	//Update(user *Employee) error
	//Store(user *Employee) (entity.ID, error)
}