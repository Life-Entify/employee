package employee

import (
	"github.com/graphql-go/graphql"
	person "github.com/life-entify/person/v1/graph/schemas"
)

var EmployeeKeywordInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "EmployeeKeywordInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"employee_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"department_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.String),
		},
	},
})
var LoginInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "LoginInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"username": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"department_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"password": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
var EmployeeInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "EmployeeInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"status": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"employee_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"department_ids": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.String),
		},
		"logins": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(LoginInputType),
		},
		"person": &graphql.InputObjectFieldConfig{
			Type: person.PersonInputType,
		},
	},
})
