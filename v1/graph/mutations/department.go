package employee

import (
	"github.com/graphql-go/graphql"
	dept_schemas "github.com/life-entify/employee/v1/graph/schemas/department"
)

func CreateDepartment(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Create Department",
		Type:        dept_schemas.DepartmentType,
		Args: graphql.FieldConfigArgument{
			"department": &graphql.ArgumentConfig{
				Type: dept_schemas.DepartmentInputType,
			},
		},
		Resolve: resolver,
	}
}
func UpdateDepartment(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Update Department",
		Type:        dept_schemas.DepartmentType,
		Args: graphql.FieldConfigArgument{
			"_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"department": &graphql.ArgumentConfig{
				Type: dept_schemas.DepartmentInputType,
			},
		},
		Resolve: resolver,
	}
}
func DeleteDepartment(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Delete Department",
		Type:        dept_schemas.DepartmentType,
		Args: graphql.FieldConfigArgument{
			"_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: resolver,
	}
}
