package employee

import (
	"context"
	"reflect"

	common "github.com/life-entify/employee/common"
	"github.com/life-entify/employee/errors"
	"github.com/life-entify/employee/v1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	COLLECTION = "employees"
)

func (db *MongoDB) ConnectEmp() (*mongo.Client, *mongo.Collection) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.uri))
	if err != nil {
		panic(err)
	}
	collection := client.Database(db.database).Collection(COLLECTION, &options.CollectionOptions{})
	return client, collection
}

func (db *MongoDB) GetNextEmployeeId(ctx context.Context) (int64, error) {
	filter := bson.D{}
	client, collection := db.ConnectEmp()
	defer MongoDisconnect(client)
	opts := options.Find().SetSort(bson.D{{Key: "employee_id", Value: -1}}).SetLimit(1)
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return 0, err
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	if len(results) > 0 {
		var resultEmployee employee.Employee
		common.ToJSONStruct(results[0], &resultEmployee)
		return resultEmployee.EmployeeId + 1, nil
	}
	return 1, nil
}

func (db *MongoDB) UpdateEmployee(ctx context.Context, _id primitive.ObjectID, p *employee.Employee) (*mongo.UpdateResult, error) {
	client, coll := db.ConnectEmp()
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
func (db *MongoDB) CreateEmployee(ctx context.Context, pt *employee.Employee) (*employee.Employee, error) {
	client, coll := db.ConnectEmp()
	defer MongoDisconnect(client)
	employeeId, err := db.GetNextEmployeeId(ctx)
	if err != nil {
		return nil, err
	}
	pt.EmployeeId = employeeId
	var jsonEmployee interface{}
	err = common.ToJSONStruct(pt, &jsonEmployee)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	value, err := coll.InsertOne(ctx, &jsonEmployee)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	if oid, ok := value.InsertedID.(primitive.ObjectID); ok {
		pt.XId = oid.Hex()
	}
	return pt, nil
}
func (db *MongoDB) FindEmployeeById(ctx context.Context, id primitive.ObjectID) (*employee.Employee, error) {
	client, coll := db.ConnectEmp()
	defer MongoDisconnect(client)
	filter := bson.M{"_id": id}

	var (
		result      bson.M
		newEmployee employee.Employee
	)
	err := coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	common.ToJSONStruct(result, &newEmployee)
	return &newEmployee, nil
}
func (db *MongoDB) FindEmployeesByPersonId(ctx context.Context, ids []int64) ([]*employee.Employee, error) {
	client, coll := db.ConnectEmp()
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
	var resultEmployees []*employee.Employee
	for _, pt := range results {
		var resultEmployee employee.Employee
		common.ToJSONStruct(pt, &resultEmployee)
		resultEmployees = append(resultEmployees, &resultEmployee)
	}
	return resultEmployees, nil
}
func (db *MongoDB) FindEmployees(ctx context.Context, filterObj *employee.Employee, page *Pagination) ([]*employee.Employee, error) {
	client, coll := db.ConnectEmp()
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

		if filterObj.EmployeeId != 0 {
			filterItems = append(filterItems, bson.M{"employee_id": filterObj.EmployeeId})
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
	var resultEmployees []*employee.Employee
	for _, pt := range results {
		var resultEmployee employee.Employee
		common.ToJSONStruct(pt, &resultEmployee)
		resultEmployees = append(resultEmployees, &resultEmployee)
	}
	return resultEmployees, nil
}
