package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gozero_example/server/internal/config"
	"gozero_example/server/internal/middleware"
)

type ServiceContext struct {
	Config  config.Config
	Auth    rest.Middleware
	CallLog rest.Middleware
	DB      *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DB.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:  c,
		Auth:    middleware.NewAuthMiddleware(db, c).Handle,
		CallLog: middleware.NewCallLogMiddleware().Handle,
		DB:      db,
	}
}
