package ascon

type Bit bool

func (b Bit) String() string {
	if b {
		return "1"
	}
	return "0"
}
