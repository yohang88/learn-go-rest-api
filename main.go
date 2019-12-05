package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"strconv"
)

type Employee struct {
	Id   int	`json:"id"`
	Name string	`json:"name"`
	City string	`json:"city"`
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())

	// Routes
	e.GET("/employees", getEmployees)
	e.GET("/employees/:id", getEmployee)
	e.POST("/employees", createEmployee)

	e.Logger.Fatal(e.Start(":8000"))
}

func getEmployees(c echo.Context) error {
	var employee Employee
	var employees []Employee

	db := connect()
	defer db.Close()

	rows, _ := db.Query(`SELECT * FROM employees`)

	for rows.Next() {
		err := rows.Scan(&employee.Id, &employee.Name, &employee.City)

		if err != nil {
			log.Fatal(err)
		}

		employees = append(employees, employee)
	}

	return c.JSON(http.StatusOK, employees)
}

func getEmployee(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var employee Employee

	db := connect()
	defer db.Close()

	row := db.QueryRow(`SELECT * FROM employees WHERE id = ?`, id)

	err := row.Scan(&employee.Id, &employee.Name, &employee.City)

	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, employee)
}

func createEmployee(c echo.Context) error {
	input := new(Employee)
	err := c.Bind(input)

	db := connect()
	defer db.Close()

	res, err := db.Exec(`INSERT INTO employees (name, city) VALUES (?, ?)`, input.Name, input.City)

	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	var employee Employee

	row := db.QueryRow(`SELECT * FROM employees WHERE id = ?`, id)

	err = row.Scan(&employee.Id, &employee.Name, &employee.City)

	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusCreated, employee)
}