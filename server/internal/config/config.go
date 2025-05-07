package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Auth struct {
		SignatureExpireMinutes int64
	}
	DB struct {
		DSN string
	}
}
