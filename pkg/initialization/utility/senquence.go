package utility

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model/meta"
	"ceres/pkg/utility/net"
	"ceres/pkg/utility/sequence"

	"github.com/gotomicro/ego/core/elog"
)

// Snowflake init logic, have to first check the ip
// use the database id as the

// AccountSequnece which will generate the Comer UIN
var AccountSequnce sequence.Senquence

func initSequnece() (err error) {
	machineIP := net.GetDomianIP()
	machineID := meta.GetMachineID(mysql.DB, machineIP)
	elog.Infof("Machine ID is %s", machineID)

	// Create snowflake sequences 

	return
}
