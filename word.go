package ascon

type Word [64]Bit

func NewWord(v uint64) Word {
	var w Word
	for i := 0; i < 64; i++ {
		w[63-i] = v&1 > 0
		v >>= 1
	}
	return w
}

func (w *Word) RotateLeft() {
	t := w[0]
	for i := 0; i < 63; i++ {
		w[i] = w[i+1]
	}
	w[63] = t
}

func (w *Word) XOR(x Word) Word {
	var y Word
	for i := 0; i < 64; i++ {
		y[i] = w[i] != x[i]
	}
	return y
}
