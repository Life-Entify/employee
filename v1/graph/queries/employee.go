package employee

import (
	"github.com/graphql-go/graphql"
	emp_schemas "github.com/life-entify/employee/v1/graph/schemas/employee"
	schemas "github.com/life-entify/person/v1/graph/schemas"
)

func GetEmployees(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Get Employees",
		Type:        graphql.NewList(emp_schemas.EmployeeType),
		Args: graphql.FieldConfigArgument{
			"keyword": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(graphql.InputObjectConfig{
					Name: "EmpKeywordInputType",
					Fields: graphql.InputObjectConfigFieldMap{
						"employee": &graphql.InputObjectFieldConfig{
							Type: emp_schemas.EmployeeKeywordInputType,
						},
						"person": &graphql.InputObjectFieldConfig{
							Type: schemas.KeywordPersonInputType,
						},
					},
				}),
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
