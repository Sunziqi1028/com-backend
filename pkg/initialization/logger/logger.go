package logger

import "github.com/gotomicro/ego/core/elog"

// Logger elog.Logger instancd
var Logger *elog.Component

// Init init the logger
// FIXME: should complete the complte configuration of logger
func Init() error {
	Logger = elog.DefaultLogger
	return nil
}
