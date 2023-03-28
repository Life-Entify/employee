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
	DEPT_COLLECTION = "emp_depts"
)

func (db *MongoDB) ConnectDept() (*mongo.Client, *mongo.Collection) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.uri))
	if err != nil {
		panic(err)
	}
	collection := client.Database(db.database).Collection(DEPT_COLLECTION, &options.CollectionOptions{})
	return client, collection
}
func (db *MongoDB) UpdateDepartment(ctx context.Context, _id primitive.ObjectID, p *employee.Department) (*mongo.UpdateResult, error) {
	client, coll := db.ConnectDept()
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
func (db *MongoDB) DeleteDepartment(ctx context.Context, _id primitive.ObjectID) (*mongo.DeleteResult, error) {
	client, coll := db.ConnectDept()
	defer MongoDisconnect(client)
	value, err := coll.DeleteOne(ctx, bson.D{{Key: "_id", Value: _id}})
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	return value, nil
}
func (db *MongoDB) CreateDepartment(ctx context.Context, dept *employee.Department) (*employee.Department, error) {
	client, coll := db.ConnectDept()
	defer MongoDisconnect(client)
	var jsonDept interface{}
	err := common.ToJSONStruct(dept, &jsonDept)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	value, err := coll.InsertOne(ctx, &jsonDept)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	if oid, ok := value.InsertedID.(primitive.ObjectID); ok {
		dept.XId = oid.Hex()
	}
	return dept, nil
}
func (db *MongoDB) FindDepartmentById(ctx context.Context, id primitive.ObjectID) (*employee.Department, error) {
	client, coll := db.ConnectDept()
	defer MongoDisconnect(client)
	filter := bson.M{"_id": id}

	var (
		result        bson.M
		newDepartment employee.Department
	)
	err := coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	common.ToJSONStruct(result, &newDepartment)
	return &newDepartment, nil
}
func (db *MongoDB) FindDepartments(ctx context.Context, filterObj *employee.Department, page *Pagination) ([]*employee.Department, error) {
	client, coll := db.ConnectDept()
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
		if filterObj.Name != "" {
			filterItems = append(filterItems, bson.M{"name": filterObj.Name})
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
	if reflect.ValueOf(filterObj).IsZero() && len(results) == 0 {
		defaultDepts := []employee.Department{
			{
				Name:        "Records",
				Description: "Medical records department",
				App:         "records",
			},
			{
				Name:        "Accounts",
				Description: "Account Department for finance functions",
				App:         "accounts",
			},
			{
				Name:        "HR",
				Description: "Human resource management department",
				App:         "accounts",
			},
		}
		var depts []interface{}
		common.ToJSONStruct(defaultDepts, &depts)
		_, err := coll.InsertMany(ctx, depts)
		if err != nil {
			return nil, errors.Errorf(err.Error())
		}
		cursor, err = coll.Find(ctx, filter, option)
		if err != nil {
			return nil, errors.Errorf(err.Error())
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.TODO(), &results); err != nil {
			return nil, errors.Errorf(err.Error())
		}
	}
	var resultDepartments []*employee.Department
	common.ToJSONStruct(results, &resultDepartments)
	return resultDepartments, nil
}
