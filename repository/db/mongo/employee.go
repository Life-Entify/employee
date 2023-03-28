package employee

import (
	"context"
	"fmt"
	"reflect"
	"strings"

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
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
func (db *MongoDB) UpdateEmployee(ctx context.Context, _id primitive.ObjectID, emp *employee.Employee) (*mongo.UpdateResult, error) {
	client, coll := db.ConnectEmp()
	defer MongoDisconnect(client)
	filter := bson.D{primitive.E{Key: "_id", Value: _id}}
	upsert := true
	opts := options.UpdateOptions{
		Upsert: &upsert,
	}
	if !reflect.ValueOf(emp.Logins).IsZero() && len(emp.Logins) > 0 {
		opts.ArrayFilters = &options.ArrayFilters{
			Filters: bson.A{bson.M{"x.department_id": emp.Logins[0].DepartmentId}},
		}
	}
	var (
		setUpdate    = bson.M{}
		onDateInsert = bson.M{}
	)

	values := reflect.ValueOf(emp).Elem()
	fields := values.Type()
	arrayFields := []string{"logins"}
	var err error
	for i := 0; i < values.NumField(); i++ {
		fieldValue := values.Field(i)
		fieldTag := strings.Split(fields.Field(i).Tag.Get("json"), ",")[0]
		if contains(arrayFields, fieldTag) {
			for j := 0; j < fieldValue.Len(); j++ {
				arrayObject := fieldValue.Index(j).Elem()
				arrayObjectType := arrayObject.Type() //struct metadata
				for k := 0; k < arrayObject.NumField(); k++ {
					arrayObjectValue := arrayObject.Field(k)
					arrayObjectField := arrayObjectType.Field(k)
					if arrayObjectValue.CanInterface() {
						value := arrayObjectValue.Interface()
						tag := strings.Split(arrayObjectField.Tag.Get("json"), ",")[0]
						// check if the value if the whole struct is inital value
						if reflect.Zero(arrayObjectType).Interface() != value {
							if tag == "password" {
								value, err = common.HashPassword(value.(string))
								if err != nil {
									return nil, errors.Errorf(err.Error())
								}
							}
							setUpdate[fmt.Sprintf("%s.$[x]."+tag, fieldTag)] = value
							onDateInsert[tag] = value
						}
					}
				}
			}
		} else if fieldTag == "department_ids" {

		} else {
			if fieldValue.CanInterface() {
				if reflect.Zero(fieldValue.Type()).Interface() != fieldValue.Interface() {
					tag := strings.Split(fields.Field(i).Tag.Get("json"), ",")[0]
					setUpdate[tag] = fieldValue.Interface()
				}
			}
		}
	}
	updateData := bson.M{
		"$set": setUpdate,
	}

	value, err := coll.UpdateOne(ctx, filter, updateData, &opts)
	if err != nil {
		if strings.Contains(err.Error(), "The path 'logins' must exist") {
			_, err1 := coll.UpdateOne(ctx, filter,
				bson.M{"$set": bson.M{"logins": []interface{}{onDateInsert}}},
			)
			if err1 != nil {
				return nil, errors.Errorf("%s => caused by $s", err1.Error(), err.Error())
			}
			value, err = coll.UpdateOne(ctx, filter, updateData, &opts)
			if err != nil {
				return nil, errors.Errorf("%s => caused by %s", err.Error(), err.Error())
			}
		}
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
