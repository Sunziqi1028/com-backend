package sequence

import "sync"

const (
	timestampBits  = uint(41)
	datacenterBits = uint(5)
	workerBits     = uint(5)
	sequenceBits   = uint(12)
)

/// Snowflake algorithm implementation
type Snowflake struct {
	sync.Mutex
	epoch uint64
}
