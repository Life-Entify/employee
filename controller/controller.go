package employee

import (
	config "github.com/life-entify/employee/config"
	repo "github.com/life-entify/employee/repository"
	db "github.com/life-entify/employee/repository/db/mongo"
)

type Controller struct {
	repo.Repository
	Config config.IConfig
}

const (
	Mongo    = "MONGODB"
	MySQL    = "MYSQL"
	PostGres = "POSTGRES"
)

func NewController(config config.IConfig) *Controller {
	return &Controller{
		db.NewMongoDB(config),
		config,
	}
}
