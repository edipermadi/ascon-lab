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

func (w Word) ROL(n int) Word {
	x := w
	for j := 0; j < n; j++ {
		t := x[0]
		for i := 0; i < len(w)-1; i++ {
			x[i] = x[i+1]
		}
		x[len(w)-1] = t
	}
	return x
}

func (w Word) ROR(n int) Word {
	x := w
	for j := 0; j < n; j++ {
		t := x[len(w)-1]
		for i := 0; i < len(w)-1; i++ {
			x[i+1] = x[i]
		}
		x[0] = t
	}
	return x
}

func (w Word) XOR(x Word) Word {
	var y Word
	for i := 0; i < len(w); i++ {
		y[i] = w[i] != x[i]
	}
	return y
}

func (w Word) AND(x Word) Word {
	var y Word
	for i := 0; i < len(w); i++ {
		y[i] = w[i] && x[i]
	}
	return y
}

func (w Word) Not() Word {
	var y Word
	for i := 0; i < len(w); i++ {
		y[i] = !w[i]
	}
	return y
}

func (w Word) UInt() uint64 {
	var y uint64
	for i := 0; i < len(w); i++ {
		y = y << 1
		if w[i] {
			y = y | 1
		}
	}
	return y
}
