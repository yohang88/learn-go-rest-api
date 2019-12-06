package entities

type Payroll struct {
	Id     int	`json:"id"`
	Year   string	`json:"year"`
	Month  string	`json:"month"`
	Salary string	`json:"salary"`
}