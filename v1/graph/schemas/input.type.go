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
			Type: graphql.String,
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
			Type: graphql.String,
		},
		"old_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"person": &graphql.InputObjectFieldConfig{
			Type: person.PersonInputType,
		},
		"next_of_kins": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(NextOfKinInputType),
		},
	},
})
var NextOfKinInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "NextOfKinInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"relationship": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"next_of_kin": &graphql.InputObjectFieldConfig{
			Type: person.ProfileInputType,
		},
	},
})
var NextOfKinMetaInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "NextOfKinMetaInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"relationship": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"person_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
