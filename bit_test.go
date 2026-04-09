package ascon_test

import (
	"fmt"
	"testing"

	"github.com/edipermadi/ascon-lab"
	"github.com/stretchr/testify/assert"
)

func TestBit_String(t *testing.T) {
	s := ascon.Bit(true)
	c := ascon.Bit(false)
	assert.Equal(t, "1", fmt.Sprintf("%v", s))
	assert.Equal(t, "0", fmt.Sprintf("%v", c))
}
