package ascon_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/edipermadi/ascon-lab"
	"github.com/stretchr/testify/assert"
)

func TestNewWord(t *testing.T) {
	type testCase struct {
		Given    uint64
		Expected ascon.Word
	}

	H := ascon.Bit(true)
	L := ascon.Bit(false)

	testCases := []testCase{
		{
			Given: 0x0000000000000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000000000001,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H},
		},
		{
			Given: 0x0000000000000002,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L},
		},
		{
			Given: 0x0000000000000004,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L},
		},
		{
			Given: 0x0000000000000008,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L},
		},
		{
			Given: 0x0000000000000010,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L},
		},
		{
			Given: 0x0000000000000020,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L},
		},
		{
			Given: 0x0000000000000040,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000000000080,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000000000100,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000000000200,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000000000400,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000000000800,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000000001000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000000002000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000000004000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000000008000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000000010000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000000020000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000000040000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000000080000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000000100000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000000200000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000000400000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000000800000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000001000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000002000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000004000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000008000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000010000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000020000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000040000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000080000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000100000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000200000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000400000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000000800000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000001000000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000002000000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000004000000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000008000000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000010000000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000020000000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000040000000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000080000000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000100000000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000200000000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000400000000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0000800000000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0001000000000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0002000000000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0004000000000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0008000000000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0010000000000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0020000000000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0040000000000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0080000000000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0100000000000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0200000000000000,
			Expected: ascon.Word{
				L, L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0400000000000000,
			Expected: ascon.Word{
				L, L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x0800000000000000,
			Expected: ascon.Word{
				L, L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x1000000000000000,
			Expected: ascon.Word{
				L, L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x2000000000000000,
			Expected: ascon.Word{
				L, L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x4000000000000000,
			Expected: ascon.Word{
				L, H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
		{
			Given: 0x8000000000000000,
			Expected: ascon.Word{
				H, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L,
				L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L, L},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Given_%s", strconv.FormatUint(tc.Given, 16)), func(t *testing.T) {
			w := ascon.NewWord(tc.Given)
			assert.Equal(t, tc.Expected, w)
			assert.Equal(t, tc.Given, w.UInt(), "expected %v actual %v", tc.Given, w.UInt())
		})
	}
}

func TestWord_XOR(t *testing.T) {
}
