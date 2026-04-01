package ascon

func BasicSBOX(in uint) uint {
	x0 := ((in >> 4) & 1) > 0
	x1 := ((in >> 3) & 1) > 0
	x2 := ((in >> 2) & 1) > 0
	x3 := ((in >> 1) & 1) > 0
	x4 := ((in >> 0) & 1) > 0

	y0 := xor(and(x4, x1), x3, and(x2, x1), x2, and(x1, x0), x1, x0)
	y1 := xor(x4, and(x3, x2), and(x3, x1), x3, and(x2, x1), x2, x1, x0)
	y2 := xor(and(x4, x3), x4, x2, x1, true)
	y3 := xor(and(x4, x0), x4, and(x3, x0), x3, x2, x1, x0)
	y4 := xor(and(x4, x1), x4, x3, and(x1, x0), x1)

	out := uint(0)
	if y0 {
		out += 16
	}
	if y1 {
		out += 8
	}
	if y2 {
		out += 4
	}
	if y3 {
		out += 2
	}
	if y4 {
		out += 1
	}
	return out
}

func FactoredSBOX(in uint) uint {
	x0 := ((in >> 4) & 1) > 0
	x1 := ((in >> 3) & 1) > 0
	x2 := ((in >> 2) & 1) > 0
	x3 := ((in >> 1) & 1) > 0
	x4 := ((in >> 0) & 1) > 0

	a0 := xor(x4, x0)
	a1 := xor(x2, x1)
	a2 := xor(x4, x3)

	b0 := and(x1, xor(a0, true))
	b1 := and(a1, xor(x1, true))
	b2 := and(x3, xor(a1, true))
	b3 := and(a2, xor(x3, true))
	b4 := and(a0, xor(a2, true))

	c0 := xor(a0, b1)
	c1 := xor(x1, b2)
	c2 := xor(a1, b3)
	c3 := xor(x3, b4)
	c4 := xor(a2, b0)

	y0 := xor(c0, c4)
	y1 := xor(c0, c1)
	y2 := xor(c2, true)
	y3 := xor(c3, c2)
	y4 := c4

	out := uint(0)
	if y0 {
		out += 16
	}
	if y1 {
		out += 8
	}
	if y2 {
		out += 4
	}
	if y3 {
		out += 2
	}
	if y4 {
		out += 1
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
	return cnt&1 > 0
}
