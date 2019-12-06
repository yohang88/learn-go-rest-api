package mysql

import (
	"database/sql"
	"employees/entities"
	"employees/repositories"
	"log"
)

type EmployeeRepositoryMysql struct {
	DB *sql.DB
}

func NewEmployeeRepositoryMysql(s *sql.DB) repositories.EmployeeRepository {
	return &EmployeeRepositoryMysql{DB: s}
}

func (e EmployeeRepositoryMysql) FindAll() ([]entities.Employee, error) {
	var employee entities.Employee
	var employees []entities.Employee

	rows, _ := e.DB.Query(`SELECT * FROM employees`)

	for rows.Next() {
		err := rows.Scan(&employee.Id, &employee.Name, &employee.City)

		if err != nil {
			log.Fatal(err)
		}

		employees = append(employees, employee)
	}

	return employees, nil
}


func (e EmployeeRepositoryMysql) Find(id int) (entities.Employee, error) {
	var employee entities.Employee

	row := e.DB.QueryRow(`SELECT * FROM employees WHERE id = ?`, id)

	err := row.Scan(&employee.Id, &employee.Name, &employee.City)

	if err != nil && err == sql.ErrNoRows {
		return employee, err
	}

	if err != nil {
		log.Fatal(err)
	}

	return employee, nil
}

func (e EmployeeRepositoryMysql) Store(input *entities.Employee) (entities.Employee, error) {
	res, err := e.DB.Exec(`INSERT INTO employees (name, city) VALUES (?, ?)`, input.Name, input.City)

	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	var employee entities.Employee

	row := e.DB.QueryRow(`SELECT * FROM employees WHERE id = ?`, id)

	err = row.Scan(&employee.Id, &employee.Name, &employee.City)

	return employee, err
}

func (e EmployeeRepositoryMysql) Update(id int, input *entities.Employee) (entities.Employee, error) {
	employee, _ := e.Find(id)

	employee.Name = input.Name
	employee.City = input.City

	_, err := e.DB.Exec(`UPDATE employees SET name = ?, city = ? WHERE id = ?`, employee.Name, employee.City, id)

	if err != nil {
		log.Fatal(err)
	}

	employee, _ = e.Find(id)

	return employee, err
}

func (e EmployeeRepositoryMysql) Destroy(id int) error {
	employee, _ := e.Find(id)

	_, err := e.DB.Exec(`DELETE FROM employees WHERE id = ?`, employee.Id)

	if err != nil {
		log.Fatal(err)
	}

	return err
}
