package employee

import (
	"github.com/graphql-go/graphql"
	dept_schemas "github.com/life-entify/employee/v1/graph/schemas/department"
)

func GetDepartments(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Get Departments",
		Type:        graphql.NewList(dept_schemas.DepartmentType),
		Args: graphql.FieldConfigArgument{
			"keyword": &graphql.ArgumentConfig{
				Type: dept_schemas.DepartmentInputType,
			},
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"skip": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: resolver,
	}
}
