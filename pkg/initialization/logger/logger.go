package logger

import "github.com/gotomicro/ego/core/elog"

// Logger elog.Logger instancd
var Logger *elog.Component

// Init init the logger
func Init() error {
	Logger = elog.Load("ceres.logger").Build()
	elog.DefaultLogger = Logger

	return nil
}
