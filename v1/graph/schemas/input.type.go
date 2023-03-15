package patient

import (
	"github.com/graphql-go/graphql"
	person "github.com/life-entify/person/v1/graph/schemas"
)

var PatientKeywordInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PatientKeywordInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"patient_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"old_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
var PatientInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PatientInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"patient_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"old_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"person": &graphql.InputObjectFieldConfig{
			Type: person.PersonInputType,
		},
	},
})
