package svc

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gozero_example/admin/internal/config"
	"gozero_example/model"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DB.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.User{}, &model.Credential{})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
