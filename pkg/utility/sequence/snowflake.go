package sequence

import (
	"crypto/rand"
	"math/big"
	"sync"
	"time"
)

const (
	timestampBits   = uint(41)
	workerBits      = uint(5)
	sequenceBits    = uint(12)
	sequenceMask    = uint64(-1 ^ (-1 << sequenceBits))
	workeridShift   = sequenceBits
	datacenterShift = sequenceBits + workerBits
)

// Snowflake algorithm implementation
type Snowflake struct {
	sync.Mutex
	epoch        uint64
	timestamp    uint64
	workerID     uint64
	detacenterID uint64
	sequence     uint64
}

// NewSnowflake create a new Snowflake algorithm
func NewSnowflake(epoch, workerID uint64) (snowflake *Snowflake) {
	centerID, _ := rand.Int(rand.Reader, big.NewInt(int64(10)))
	return &Snowflake{
		epoch:        epoch,
		workerID:     workerID,
		detacenterID: uint64(centerID.Bit(32)),
	}
}

// Next implementa the Sequence interface
func (flake *Snowflake) Next() (seq uint64) {
	flake.Lock()
	now := time.Now().UnixNano() / 1000000
	if flake.timestamp == uint64(now) {
		flake.sequence = (flake.sequence + 1) & uint64(sequenceMask)
		if flake.sequence == 0 {
			for now <= int64(flake.timestamp) {
				now = time.Now().UnixNano() / 1000000
			}
		}

	} else {
		flake.sequence = 0
	}
	t := now - int64(flake.epoch)
	flake.timestamp = uint64(now)
	seq = uint64(
		(t << int64(timestampBits)) |
			(int64(flake.detacenterID) << int64(datacenterShift)) |
			(int64(flake.workerID) << int64(workeridShift)) |
			int64(flake.sequence))
	flake.Unlock()
	return
}
