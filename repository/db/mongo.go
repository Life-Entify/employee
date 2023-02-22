package patient

import (
	"context"
	"reflect"

	common "github.com/life-entify/patient/common"
	config "github.com/life-entify/patient/config"
	"github.com/life-entify/patient/errors"
	"github.com/life-entify/patient/v1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	COLLECTION = "patients"
)

type MongoDB struct {
	uri      string
	database string
}

func NewMongoDB(config config.IConfig) *MongoDB {
	url := config.GetDBUrl()
	return &MongoDB{
		uri:      url,
		database: config.GetDBName(),
	}
}
func (db *MongoDB) Connect() (*mongo.Client, *mongo.Collection) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.uri))
	if err != nil {
		panic(err)
	}
	collection := client.Database(db.database).Collection(COLLECTION, &options.CollectionOptions{})
	return client, collection
}
func MongoDisconnect(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
func (db *MongoDB) GetNextPatientId(ctx context.Context) (int64, error) {
	filter := bson.D{}
	client, collection := db.Connect()
	defer MongoDisconnect(client)
	opts := options.Find().SetSort(bson.D{{Key: "patient_id", Value: -1}}).SetLimit(1)
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return 0, err
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	if len(results) > 0 {
		var resultPatient patient.Patient
		common.ToJSONStruct(results[0], &resultPatient)
		return resultPatient.PatientId + 1, nil
	}
	return 1, nil
}

func (db *MongoDB) UpdatePatient(ctx context.Context, _id primitive.ObjectID, p *patient.Patient) (*mongo.UpdateResult, error) {
	client, coll := db.Connect()
	defer MongoDisconnect(client)
	var update bson.M
	err := common.ToJSONStruct(p, &update)
	if err != nil {
		return nil, err
	}
	value, err := coll.UpdateOne(ctx, bson.D{{Key: "_id", Value: _id}}, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	return value, nil
}
func (db *MongoDB) CreatePatient(ctx context.Context, pt *patient.Patient) (*patient.Patient, error) {
	client, coll := db.Connect()
	defer MongoDisconnect(client)
	patientId, err := db.GetNextPatientId(ctx)
	if err != nil {
		return nil, err
	}
	pt.PatientId = patientId
	var jsonPatient interface{}
	err = common.ToJSONStruct(pt, &jsonPatient)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	value, err := coll.InsertOne(ctx, &jsonPatient)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	if oid, ok := value.InsertedID.(primitive.ObjectID); ok {
		pt.XId = oid.Hex()
	}
	return pt, nil
}
func (db *MongoDB) FindPatientById(ctx context.Context, id primitive.ObjectID) (*patient.Patient, error) {
	client, coll := db.Connect()
	defer MongoDisconnect(client)
	filter := bson.M{"_id": id}

	var (
		result     bson.M
		newPatient patient.Patient
	)
	err := coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	common.ToJSONStruct(result, &newPatient)
	return &newPatient, nil
}
func (db *MongoDB) FindPatientsByPersonId(ctx context.Context, ids []string) ([]*patient.Patient, error) {
	client, coll := db.Connect()
	defer MongoDisconnect(client)
	opts := options.Find().SetSort(bson.D{{Key: "person_id", Value: 1}})
	filter := bson.D{{Key: "person_id", Value: bson.D{{Key: "$in", Value: ids}}}}

	cursor, err := coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, errors.Errorf(err.Error())
	}
	var resultPatients []*patient.Patient
	for _, pt := range results {
		var resultPatient patient.Patient
		common.ToJSONStruct(pt, &resultPatient)
		resultPatients = append(resultPatients, &resultPatient)
	}
	return resultPatients, nil
}
func (db *MongoDB) FindPatients(ctx context.Context, filterObj *patient.Patient, page *Pagination) ([]*patient.Patient, error) {
	client, coll := db.Connect()
	defer MongoDisconnect(client)
	option := options.Find().SetSort(bson.D{{Key: "_id", Value: 1}})
	if !reflect.ValueOf(page).IsZero() {
		if page.Skip != 0 {
			option.SetSkip(page.Skip)
		}
		if page.Limit != 0 {
			option.SetLimit(page.Limit)
		}
	}
	var filter = bson.M{}
	if !reflect.ValueOf(filterObj).IsZero() {
		filterItems := []bson.M{}
		if filterObj.XId != "" {
			id, err := primitive.ObjectIDFromHex(filterObj.XId)
			if err != nil {
				return nil, nil
			}
			filterItems = append(filterItems, bson.M{"_id": id})
		}
		if filterObj.OldId != "" {
			filterItems = append(filterItems, bson.M{"old_id": filterObj.OldId})
		}
		if filterObj.PatientId != 0 {
			filterItems = append(filterItems, bson.M{"patient_id": filterObj.PatientId})
		}
		if len(filterItems) > 0 {
			filter["$or"] = filterItems
		}
	}
	cursor, err := coll.Find(ctx, filter, option)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	defer cursor.Close(ctx)
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, errors.Errorf(err.Error())
	}
	var resultPatients []*patient.Patient
	for _, pt := range results {
		var resultPatient patient.Patient
		common.ToJSONStruct(pt, &resultPatient)
		resultPatients = append(resultPatients, &resultPatient)
	}
	return resultPatients, nil
}
