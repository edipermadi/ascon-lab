package ascon

func Spread(v uint32) uint64 {
	x := uint64(v) // Promote to 64-bit
	// Spread 16 bits to 32 bits (gaps of 16)
	x = (x | (x << 16)) & 0x0000FFFF0000FFFF
	// Spread 8 bits to 16 bits (gaps of 8)
	x = (x | (x << 8)) & 0x00FF00FF00FF00FF
	// Spread 4 bits to 8 bits (gaps of 4)
	x = (x | (x << 4)) & 0x0F0F0F0F0F0F0F0F
	// Spread 2 bits to 4 bits (gaps of 2)
	x = (x | (x << 2)) & 0x3333333333333333
	// Spread 1 bit to 2 bits (gaps of 1)
	x = (x | (x << 1)) & 0x5555555555555555

	// At this point, x has the bits in positions 0,2,4...
	// To make odd bit == even bit, we can OR the result with a shift
	return x | (x << 1)
}
