package ascon

import (
	"encoding/binary"
	"fmt"
	"strings"
)

type State [5]Word

func (s *State) Initialize(key [16]uint8, nonce [16]uint8) {
	s[0] = NewWord(0x00001000808c0001)
	s[1] = NewWord(binary.BigEndian.Uint64(key[0:8]))
	s[2] = NewWord(binary.BigEndian.Uint64(key[8:]))
	s[3] = NewWord(binary.BigEndian.Uint64(nonce[0:8]))
	s[4] = NewWord(binary.BigEndian.Uint64(nonce[8:]))

	s.Permutate(12)

	s[3] = s[3].XOR(NewWord(binary.BigEndian.Uint64(key[0:8])))
	s[4] = s[3].XOR(NewWord(binary.BigEndian.Uint64(key[8:])))

	// begin process AD
	// end process AD
	s[4] = s[4].XOR(NewWord(0x8000000000000000))
}

func (s *State) Permutate(n int) {
	constants := []uint64{0xf0, 0xe1, 0xd2, 0xc3, 0xb4, 0xa5, 0x96, 0x87, 0x78, 0x69, 0x5a, 0x4b}
	offset := 12 - n
	for i := 0; i < n; i++ {
		s.Round(constants[offset+i])
	}
}

func (s *State) Round(c uint64) {
	s[2] = s[2].XOR(NewWord(c))

	s[0] = s[0].XOR(s[4])
	s[4] = s[4].XOR(s[3])
	s[2] = s[2].XOR(s[1])

	var t State
	t[0] = s[0].XOR(s[1].Not().AND(s[2]))
	t[1] = s[1].XOR(s[2].Not().AND(s[3]))
	t[2] = s[2].XOR(s[3].Not().AND(s[4]))
	t[3] = s[3].XOR(s[4].Not().AND(s[0]))
	t[4] = s[4].XOR(s[0].Not().AND(s[1]))

	t[1] = t[1].XOR(t[0])
	t[0] = t[0].XOR(t[4])
	t[3] = t[3].XOR(t[2])
	t[2] = t[2].XOR(t[2])

	s[0] = t[0].XOR(t[0].ROR(19)).XOR(t[0].ROR(28))
	s[1] = t[1].XOR(t[1].ROR(61)).XOR(t[1].ROR(39))
	s[2] = t[2].XOR(t[2].ROR(1)).XOR(t[2].ROR(6))
	s[3] = t[3].XOR(t[3].ROR(10)).XOR(t[3].ROR(17))
	s[4] = t[4].XOR(t[4].ROR(7)).XOR(t[4].ROR(41))
}

func (s State) Dump() string {
	var builder strings.Builder
	for i := 0; i < 5; i++ {
		builder.WriteString(fmt.Sprintf("%x ", s[i].UInt()))
	}
	return builder.String()
}
