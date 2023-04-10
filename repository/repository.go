package employee

import (
	"context"

	db "github.com/life-entify/employee/repository/db/mongo"
	"github.com/life-entify/employee/v1"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	ConnectDept() (*mongo.Client, *mongo.Collection)
	ConnectEmp() (*mongo.Client, *mongo.Collection)

	DeleteDepartment(ctx context.Context, _id primitive.ObjectID) (*mongo.DeleteResult, error)
	FindDepartmentById(ctx context.Context, id primitive.ObjectID) (*employee.Department, error)
	FindDepartments(ctx context.Context, filterObj *employee.Department, page *db.Pagination) ([]*employee.Department, error)
	UpdateDepartment(ctx context.Context, _id primitive.ObjectID, p *employee.Department) (*mongo.UpdateResult, error)
	CreateDepartment(ctx context.Context, dept *employee.Department) (*employee.Department, error)

	AddEmpDepartment(ctx context.Context, _id primitive.ObjectID, deptId string, login *employee.Login) (*employee.Employee, error)
	DeleteEmpDepartment(ctx context.Context, _id primitive.ObjectID, deptId string) (*employee.Employee, error)
	UpdateEmployee(ctx context.Context, _id primitive.ObjectID, p *employee.Employee) (*mongo.UpdateResult, error)
	CreateEmployee(ctx context.Context, employee *employee.Employee) (*employee.Employee, error)
	FindEmployeeById(ctx context.Context, id primitive.ObjectID) (*employee.Employee, error)
	FindEmployeesByEmployeeId(ctx context.Context, ids []int64) ([]*employee.Employee, error)
	FindEmployeesByPersonId(ctx context.Context, ids []int64) ([]*employee.Employee, error)
	FindEmployees(ctx context.Context, filterObj *employee.Employee, page *db.Pagination) ([]*employee.Employee, error)
}
