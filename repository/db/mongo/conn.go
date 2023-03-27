package employee

import (
	"context"

	config "github.com/life-entify/employee/config"
	"go.mongodb.org/mongo-driver/mongo"
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

func MongoDisconnect(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
