package main

import (
	"database/sql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"strconv"
)

var db *sql.DB

type Employee struct {
	Id   int	`json:"id"`
	Name string	`json:"name"`
	City string	`json:"city"`
}

func main() {
	db = connect()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"version": "1.0.0"})
	})

	// Routes
	e.GET("/employees", getEmployees)
	e.GET("/employees/:id", getEmployee)
	e.POST("/employees", createEmployee)
	e.PUT("/employees/:id", updateEmployee)
	e.DELETE("/employees/:id", deleteEmployee)

	e.Logger.Fatal(e.Start(":8000"))
}

func getEmployees(c echo.Context) error {
	var employee Employee
	var employees []Employee

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

	row := db.QueryRow(`SELECT * FROM employees WHERE id = ?`, id)

	err := row.Scan(&employee.Id, &employee.Name, &employee.City)

	if err != nil && err == sql.ErrNoRows {
		return c.NoContent(http.StatusNotFound)
	}

	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, employee)
}

func createEmployee(c echo.Context) error {
	input := new(Employee)
	err := c.Bind(input)

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


func updateEmployee(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var employee Employee

	// Check Existence
	row := db.QueryRow(`SELECT * FROM employees WHERE id = ?`, id)

	err := row.Scan(&employee.Id, &employee.Name, &employee.City)

	if err != nil && err == sql.ErrNoRows {
		return c.NoContent(http.StatusNotFound)
	}

	if err != nil {
		log.Fatal(err)
	}

	// Update if found
	input := new(Employee)
	err = c.Bind(input)

	employee.Name = input.Name
	employee.City = input.City

	_, err = db.Exec(`UPDATE employees SET name = ?, city = ? WHERE id = ?`, employee.Name, employee.City, id)

	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, employee)
}

func deleteEmployee(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// Check Existence
	row := db.QueryRow(`SELECT * FROM employees WHERE id = ?`, id)

	err := row.Scan()

	if err != nil && err == sql.ErrNoRows {
		return c.NoContent(http.StatusNotFound)
	}

	// Delete if found
	_, err = db.Exec(`DELETE FROM employees WHERE id = ?`, id)

	if err != nil {
		log.Fatal(err)
	}

	return c.NoContent(http.StatusNoContent)
}

