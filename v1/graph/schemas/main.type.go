package patient

import (
	"github.com/graphql-go/graphql"
	person "github.com/life-entify/person/v1/graph/schemas"
)

var NextOfKinType = graphql.NewObject(graphql.ObjectConfig{
	Name: "NextOfKin",
	Fields: graphql.Fields{
		"person_id": &graphql.Field{
			Type: graphql.String,
		},
		"relationship": &graphql.Field{
			Type: graphql.String,
		},
	},
})
var PatientType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Patient",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
		},
		"patient_id": &graphql.Field{
			Type: graphql.String,
		},
		"old_id": &graphql.Field{
			Type: graphql.String,
		},
		"person": &graphql.Field{
			Type: person.PersonType,
		},
		"next_of_kins": &graphql.Field{
			Type: graphql.NewList(NextOfKinType),
		},
	},
})
