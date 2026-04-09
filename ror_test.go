package ascon_test

import (
	"fmt"
	"testing"

	"github.com/edipermadi/ascon-lab"
)

func makeRange(count int) []string {
	result := make([]string, count)
	for i := 0; i < count; i++ {
		result[i] = fmt.Sprintf("%02d", i)
	}
	return result
}

func TestROR(t *testing.T) {
	x1 := makeRange(64)
	x2 := ascon.ROR(makeRange(64), 19)
	x3 := ascon.ROR(makeRange(64), 28)
	t.Logf("x1 = %v", x1)
	t.Logf("x2 = %v", x2)
	t.Logf("x3 = %v", x3)
}
