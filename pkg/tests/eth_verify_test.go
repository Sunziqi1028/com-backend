package pkg

import (
	"encoding/hex"
	"testing"
)

func TestSimpleEncoding(t *testing.T) {

	origin := "0xE2346ffF0ae08e172B2C6384F5Aa66C42c5527D9"

	res, err := hex.DecodeString(origin)
	t.Log(err)
	t.Log(res)
}
