package employee

import (
	"github.com/graphql-go/graphql"
)

var DepartmentInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "DepartmentInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"app": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
