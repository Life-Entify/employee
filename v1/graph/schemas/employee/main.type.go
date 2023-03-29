package employee

import (
	"github.com/graphql-go/graphql"
	person "github.com/life-entify/person/v1/graph/schemas"
)

var LoginType = graphql.NewObject(graphql.ObjectConfig{
	Name: "LoginType",
	Fields: graphql.Fields{
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"department_id": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
	},
})
var EmployeeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Employee",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"employee_id": &graphql.Field{
			Type: graphql.Int,
		},
		"department_ids": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"logins": &graphql.Field{
			Type: graphql.NewList(LoginType),
		},
		"person": &graphql.Field{
			Type: person.PersonType,
		},
	},
})
