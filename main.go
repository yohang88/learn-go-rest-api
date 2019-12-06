package main

import (
	"database/sql"
	"employees/entities"
	_ "github.com/go-sql-driver/mysql"

	repoMysql "employees/repositories/mysql"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"strconv"
)

var db *sql.DB

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
	employeeRepo := repoMysql.NewEmployeeRepositoryMysql(db)

	employees, _ := employeeRepo.FindAll()

	return c.JSON(http.StatusOK, employees)
}

func getEmployee(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	employeeRepo := repoMysql.NewEmployeeRepositoryMysql(db)

	employee, err := employeeRepo.Find(id)

	if err != nil && err == sql.ErrNoRows {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, employee)
}

func createEmployee(c echo.Context) error {
	input := new(entities.Employee)
	err := c.Bind(input)

	employeeRepo := repoMysql.NewEmployeeRepositoryMysql(db)

	employee, err := employeeRepo.Store(input)

	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusCreated, employee)
}


func updateEmployee(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	input := new(entities.Employee)
	err := c.Bind(input)

	employeeRepo := repoMysql.NewEmployeeRepositoryMysql(db)

	employee, err := employeeRepo.Update(id, input)

	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, employee)
}

func deleteEmployee(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	employeeRepo := repoMysql.NewEmployeeRepositoryMysql(db)

	err := employeeRepo.Destroy(id)

	if err != nil {
		log.Fatal(err)
	}

	return c.NoContent(http.StatusNoContent)
}


func connect() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/go_dev")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
