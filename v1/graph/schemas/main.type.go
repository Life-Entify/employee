package patient

import (
	"github.com/graphql-go/graphql"
	person "github.com/life-entify/person/v1/graph/schemas"
)

var PatientType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Patient",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
		},
		"patient_id": &graphql.Field{
			Type: graphql.Int,
		},
		"old_id": &graphql.Field{
			Type: graphql.String,
		},
		"person": &graphql.Field{
			Type: person.PersonType,
		},
	},
})
