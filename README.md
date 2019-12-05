# Learn Golang REST API

## Learning Path (To Do)
- Go Web Framework: Echo https://echo.labstack.com
- Basic REST API (List, View, Create, Update, Delete)
- Database Connection
- Basic Routing, Middleware, Handler, Logger

## Setup
1. Create DB, import `db/ddl.sql`
2. Configure database connection in `mysql.go`

## Run
1. `go run .`

## Build
1. `go build`
2. `./echo`

## Endpoints
### Get All Employees
```
curl -X GET http://localhost:8000/employees -H 'Accept: application/json'
```

### Get Single Employee
```
curl -X GET http://localhost:8000/employees/1 -H 'Accept: application/json'
```

### Create Employee
```
curl -X POST \
  http://localhost:8000/employees \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
	"name": "Test Name",
	"city": "Jakarta"
}'
```

### Update Employee
```
curl -X PUT \
  http://localhost:8000/employees/1 \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
	"name": "Name New",
	"city": "City New"
}'
```

### Delete Employee
```
curl -X DELETE http://localhost:8000/employees/1
```