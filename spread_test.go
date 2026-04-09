package ascon_test

import (
	"strconv"
	"testing"

	"github.com/edipermadi/ascon-lab"
)

func TestSpread(t *testing.T) {
	a := uint32(2)
	b := ascon.Spread(a)
	t.Logf("a = %v, b = %v", strconv.FormatUint(uint64(a), 16), strconv.FormatUint(b, 16))
}
