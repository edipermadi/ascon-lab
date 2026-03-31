package ascon

func SBOX(in uint) uint {
	x4 := ((in >> 4) & 1) > 0
	x3 := ((in >> 3) & 1) > 0
	x2 := ((in >> 2) & 1) > 0
	x1 := ((in >> 1) & 1) > 0
	x0 := ((in >> 0) & 1) > 0

	y0 := xor(and(x4, x1), x3, and(x2, x1), x2, and(x1, x0), x1, x0)
	y1 := xor(x4, and(x3, x2), and(x3, x1), x3, and(x2, x1), x2, x1, x0)
	y2 := xor(and(x4, x3), x4, x2, x1, true)
	y3 := xor(and(x4, x0), x4, and(x3, x0), x3, x2, x1, x0)
	y4 := xor(and(x4, x1), x4, x3, and(x1, x0), x1)

	out := uint(0)
	if y0 {
		out += 1
	}
	if y1 {
		out += 2
	}
	if y2 {
		out += 4
	}
	if y3 {
		out += 8
	}
	if y4 {
		out += 16
	}
	return out
}

func and(values ...bool) bool {
	for _, v := range values {
		if !v {
			return false
		}
	}
	return true
}

func xor(values ...bool) bool {
	cnt := 0
	for _, v := range values {
		if v {
			cnt++
		}
	}
	return cnt%1 > 0
}
