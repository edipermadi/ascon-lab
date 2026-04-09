package ascon

type Word [64]Bit

func NewWord(v uint64) Word {
	var w Word
	for i := 0; i < len(w); i++ {
		w[63-i] = v&1 > 0
		v >>= 1
	}
	return w
}

func (w *Word) RotateLeft() {
	t := w[0]
	for i := 0; i < len(w)-1; i++ {
		w[i] = w[i+1]
	}
	w[len(w)-1] = t
}

func (w *Word) XOR(x Word) Word {
	var y Word
	for i := 0; i < len(w); i++ {
		y[i] = w[i] != x[i]
	}
	return y
}

func (w *Word) UInt() uint64 {
	var y uint64
	for i := 0; i < len(w); i++ {
		if w[i] {
			y = y | 1
		}
		y = y << 1
	}
	return y
}
