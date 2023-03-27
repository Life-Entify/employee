package employee

import (
	"github.com/graphql-go/graphql"
	person "github.com/life-entify/person/v1/graph/schemas"
)

var EmployeeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Employee",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
		},
		"employee_id": &graphql.Field{
			Type: graphql.Int,
		},
		"department_ids": &graphql.Field{
			Type: graphql.NewList(graphql.Int),
		},
		"person": &graphql.Field{
			Type: person.PersonType,
		},
	},
})
