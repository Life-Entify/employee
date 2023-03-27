package employee

import (
	"github.com/graphql-go/graphql"
	emp_schemas "github.com/life-entify/employee/v1/graph/schemas/employee"
	person_schemas "github.com/life-entify/person/v1/graph/schemas"
)

func CreateEmployeeWithNok(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Create Employee With Metadata",
		Type:        emp_schemas.EmployeeType,
		Args: graphql.FieldConfigArgument{
			"profile": &graphql.ArgumentConfig{
				Type: person_schemas.ProfileInputType,
			},
			"next_of_kins": &graphql.ArgumentConfig{
				Type: graphql.NewList(person_schemas.NextOfKinMetaInputType),
			},
		},
		Resolve: resolver,
	}
}
func CreateEmployeeWithMD(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Create Employee With Metadata",
		Type:        emp_schemas.EmployeeType,
		Args: graphql.FieldConfigArgument{
			"person_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"next_of_kins": &graphql.ArgumentConfig{
				Type: graphql.NewList(person_schemas.NextOfKinMetaInputType),
			},
		},
		Resolve: resolver,
	}
}
func CreateEmployeeWithPerson(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Create Employee",
		Type:        emp_schemas.EmployeeType,
		Args: graphql.FieldConfigArgument{
			"person_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"next_of_kins": &graphql.ArgumentConfig{
				Type: graphql.NewList(person_schemas.NextOfKinInputType),
			},
		},
		Resolve: resolver,
	}
}
func CreateEmployee(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Create Employee",
		Type:        emp_schemas.EmployeeType,
		Args: graphql.FieldConfigArgument{
			"profile": &graphql.ArgumentConfig{
				Type: person_schemas.ProfileInputType,
			},
			"next_of_kins": &graphql.ArgumentConfig{
				Type: graphql.NewList(person_schemas.NextOfKinInputType),
			},
		},
		Resolve: resolver,
	}
}
func UpdateEmployeeProfile(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Update Employee",
		Type:        emp_schemas.EmployeeType,
		Args: graphql.FieldConfigArgument{
			"_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"employee": &graphql.ArgumentConfig{
				Type: emp_schemas.EmployeeInputType,
			},
			"person_xid": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"profile": &graphql.ArgumentConfig{
				Type: person_schemas.ProfileInputType,
			},
		},
		Resolve: resolver,
	}
}
