package patient

import (
	"github.com/graphql-go/graphql"
	pt_schemas "github.com/life-entify/patient/v1/graph/schemas"
	schemas "github.com/life-entify/person/v1/graph/schemas"
)

func GetPatients(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Get Patients",
		Type:        graphql.NewList(pt_schemas.PatientType),
		Args: graphql.FieldConfigArgument{
			"keyword": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(graphql.InputObjectConfig{
					Name: "KeywordInputType",
					Fields: graphql.InputObjectConfigFieldMap{
						"patient": &graphql.InputObjectFieldConfig{
							Type: pt_schemas.PatientKeywordInputType,
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
