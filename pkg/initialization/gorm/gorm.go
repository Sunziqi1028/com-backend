package gorm

import "github.com/gotomicro/ego-component/egorm"

var DB *egorm.Component

func Init() error {
	DB = egorm.Load("mysql.user").Build()
	return nil
}