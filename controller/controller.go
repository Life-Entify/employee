package patient

import (
	config "github.com/life-entify/patient/config"
	repo "github.com/life-entify/patient/repository"
	db "github.com/life-entify/patient/repository/db"
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
