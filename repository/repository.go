package patient

import (
	"context"

	db "github.com/life-entify/patient/repository/db"
	"github.com/life-entify/patient/v1"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Connect() (*mongo.Client, *mongo.Collection)
	UpdatePatient(ctx context.Context, _id primitive.ObjectID, p *patient.Patient) (*mongo.UpdateResult, error)
	CreatePatient(ctx context.Context, patient *patient.Patient) (*patient.Patient, error)
	FindPatientById(ctx context.Context, id primitive.ObjectID) (*patient.Patient, error)
	FindPatientsByPersonId(ctx context.Context, ids []int64) ([]*patient.Patient, error)
	FindPatients(ctx context.Context, filterObj *patient.Patient, page *db.Pagination) ([]*patient.Patient, error)
}
