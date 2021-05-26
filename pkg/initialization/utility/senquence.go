package utility

import (
	"ceres/pkg/utility/net"
	"ceres/pkg/utility/sequence"
	"strconv"
	"strings"

	"github.com/gotomicro/ego/core/econf"
)

// Snowflake init logic, have to first check the ip
// use the database id as the

// AccountSequnece which will generate the Comer UIN
var AccountSequnece sequence.Senquence
var BountySeqnence sequence.Senquence

func initSequnece() (err error) {
	machineIP := net.GetDomianIP()
	machineSignature := strings.Replace(machineIP, ".", "", 4)
	machineID, err := strconv.ParseInt(machineSignature, 10, 64)
	if err != nil {
		return
	}
	machineID = machineID % 32
	epoch, _ := econf.Get("ceres.snowflake.epoch").(int) //TODO: should check if this is correct
	// Create snowflake sequences
	AccountSequnece = sequence.NewSnowflake(uint64(epoch), uint64(machineID))

	return
}
