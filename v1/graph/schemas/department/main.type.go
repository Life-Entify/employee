package employee

import (
	"github.com/graphql-go/graphql"
)

var DepartmentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Department",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"app": &graphql.Field{
			Type: graphql.String,
		},
	},
})
