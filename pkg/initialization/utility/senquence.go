package utility

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model/meta"
	"ceres/pkg/utility/net"
	"ceres/pkg/utility/sequence"

	"github.com/gotomicro/ego/core/econf"
)

// Snowflake init logic, have to first check the ip
// use the database id as the

// AccountSequnece which will generate the Comer UIN
var AccountSequnece sequence.Senquence
var BountySeqnence sequence.Senquence

func initSequnece() (err error) {
	machineIP := net.GetDomianIP()
	machineID := meta.GetMachineID(mysql.DB, machineIP)
	epoch, _ := econf.Get("ceres.snowflake.epoch").(int) //TODO: should check if this is correct


	// Create snowflake sequences
	AccountSequnece = sequence.NewSnowflake(uint64(epoch), uint64(machineID))
	

	return
}
