package ascon

func Encrypt(key [16]byte, nonce [16]byte) {
	var s State
	s.Initialize(key, nonce)
}
