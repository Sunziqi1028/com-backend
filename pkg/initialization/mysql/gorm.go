package mysql

import "github.com/gotomicro/ego-component/egorm"

// DB Mysql instance with gorm
var DB *egorm.Component

// Init the mysql
func Init() error {
	DB = egorm.Load("mysql.user").Build()
	return nil
}
