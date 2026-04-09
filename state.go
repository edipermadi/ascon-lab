package ascon

import "encoding/binary"

type State [5]Word

func (s *State) Initialize(key [16]byte, nonce [16]byte) {
	s[0] = NewWord(0x00001000808c0001)
	s[1] = NewWord(binary.BigEndian.Uint64(key[0:4]))
	s[2] = NewWord(binary.BigEndian.Uint64(key[4:]))
	s[3] = NewWord(binary.BigEndian.Uint64(nonce[0:4]))
	s[4] = NewWord(binary.BigEndian.Uint64(nonce[4:]))
}

func (s *State) Permutate() {}

func (s *State) Round() {

}
