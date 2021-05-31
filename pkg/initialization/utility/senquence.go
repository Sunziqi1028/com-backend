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

// AccountSequnece generate the comer uin sequence
var AccountSequnece sequence.Senquence

// ProfileSequence generate the profile sequence
var ProfileSequence sequence.Senquence

// BountySeqnence generate the bounty sequence
var BountySeqnence sequence.Senquence

func initSequnece() (err error) {
	machineIP := net.GetDomianIP()
	machineSignature := strings.Replace(machineIP, ".", "", 4)
	machineID, err := strconv.ParseInt(machineSignature, 10, 64)
	if err != nil {
		return
	}
	machineID %= 32
	epoch, _ := econf.Get("ceres.snowflake.epoch").(int) //TODO: should check if this is correct
	// Create snowflake sequences
	AccountSequnece = sequence.NewSnowflake(uint64(epoch), uint64(machineID))
	ProfileSequence = sequence.NewSnowflake(uint64(epoch), uint64(machineID))
	BountySeqnence = sequence.NewSnowflake(uint64(epoch), uint64(machineID))
	return
}
