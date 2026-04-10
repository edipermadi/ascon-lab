package ascon

import (
	"fmt"
	"os"
)

func BasicSBOX(in uint) uint {
	x0, x1, x2, x3, x4 := unpack(in)

	y0 := xor(and(x4, x1), x3, and(x2, x1), x2, and(x1, x0), x1, x0)
	y1 := xor(x4, and(x3, x2), and(x3, x1), x3, and(x2, x1), x2, x1, x0)
	y2 := xor(and(x4, x3), x4, x2, x1, true)
	y3 := xor(and(x4, x0), x4, and(x3, x0), x3, x2, x1, x0)
	y4 := xor(and(x4, x1), x4, x3, and(x1, x0), x1)

	return pack(y0, y1, y2, y3, y4)
}

func FactoredSBOX(in uint) uint {
	x0, x1, x2, x3, x4 := unpack(in)

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

	return pack(y0, y1, y2, y3, y4)
}

func ShadowSBOX(in uint) uint {
	x0, x1, x2, x3, x4 := unpack(in)

	a0 := xor(xor(x4, x0), true)
	a1 := xor(xor(x2, x1), true)
	a2 := xor(xor(x4, x3), true)

	b0 := xor(and(a0, xor(x1, true)), true)
	b1 := xor(and(x1, xor(a1, true)), true)
	b2 := xor(and(a1, xor(x3, true)), true)
	b3 := xor(and(x3, xor(a2, true)), true)
	b4 := xor(and(a2, xor(a0, true)), true)

	c0 := xor(a0, b1, true)
	c1 := xor(x1, b2, true)
	c2 := xor(a1, b3, true)
	c3 := xor(x3, b4, true)
	c4 := xor(a2, b0, true)

	y0 := xor(c4, c0, true)
	y1 := xor(c1, c0, true)
	y2 := xor(c2, true)
	y3 := xor(c3, c2, true)
	y4 := c4

	return pack(y0, y1, y2, y3, y4)
}

func unpack(in uint) (Bit, Bit, Bit, Bit, Bit) {
	x0 := ((in >> 4) & 1) > 0
	x1 := ((in >> 3) & 1) > 0
	x2 := ((in >> 2) & 1) > 0
	x3 := ((in >> 1) & 1) > 0
	x4 := ((in >> 0) & 1) > 0
	return Bit(x0), Bit(x1), Bit(x2), Bit(x3), Bit(x4)
}

func pack(y0, y1, y2, y3, y4 Bit) uint {
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

func and(values ...Bit) Bit {
	for _, v := range values {
		if !v {
			return false
		}
	}
	return true
}

func xor(values ...Bit) Bit {
	cnt := 0
	for _, v := range values {
		if v {
			cnt++
		}
	}
	return cnt&1 > 0
}

func Simulate(in uint) uint {
	fmt.Fprintf(os.Stderr, "===== %02x\n", in)

	x0, x1, x2, x3, x4 := unpack(in)
	_x0 := !x0
	_x1 := !x1
	_x2 := !x2
	_x3 := !x3
	_x4 := !x4
	fmt.Fprintf(os.Stderr, "in = %v %v %v %v %v | %v %v %v %v %v\n", x0, x1, x2, x3, x4, _x0, _x1, _x2, _x3, _x4)

	a0 := xor(x4, x0)
	a1 := xor(x2, x1)
	a2 := xor(x4, x3)

	_a0 := xor(xor(x4, x0), true)
	_a1 := xor(xor(x2, x1), true)
	_a2 := xor(xor(x4, x3), true)

	fmt.Fprintf(os.Stderr, "a0 a1 a2 = %v %v %v | %v %v %v\n", a0, a1, a2, _a0, _a1, _a2)

	b0 := and(xor(a0, true), x1)
	b1 := and(xor(x1, true), a1)
	b2 := and(xor(a1, true), x3)
	b3 := and(xor(x3, true), a2)
	b4 := and(xor(a2, true), a0)

	_b0 := xor(and(_a0, xor(_x1, true)), true)
	_b1 := xor(and(_x1, xor(_a1, true)), true)
	_b2 := xor(and(_a1, xor(_x3, true)), true)
	_b3 := xor(and(_x3, xor(_a2, true)), true)
	_b4 := xor(and(_a2, xor(_a0, true)), true)

	fmt.Fprintf(os.Stderr, "b0 b1 b2 b3 b4 = %v %v %v %v %v | %v %v %v %v %v\n", b0, b1, b2, b3, b4, _b0, _b1, _b2, _b3, _b4)

	c0 := xor(a0, b1)
	c1 := xor(x1, b2)
	c2 := xor(a1, b3)
	c3 := xor(x3, b4)
	c4 := xor(a2, b0)

	_c0 := xor(_a0, _b1, true)
	_c1 := xor(_x1, _b2, true)
	_c2 := xor(_a1, _b3, true)
	_c3 := xor(_x3, _b4, true)
	_c4 := xor(_a2, _b0, true)

	fmt.Fprintf(os.Stderr, "c0 c1 c2 c3 c4 = %v %v %v %v %v | %v %v %v %v %v\n", c0, c1, c2, c3, c4, _c0, _c1, _c2, _c3, _c4)

	y0 := xor(c4, c0)
	y1 := xor(c1, c0)
	y2 := xor(c2, true)
	y3 := xor(c3, c2)
	y4 := c4

	_y0 := xor(_c4, _c0, true)
	_y1 := xor(_c1, _c0, true)
	_y2 := xor(_c2, true)
	_y3 := xor(_c3, _c2, true)
	_y4 := _c4

	fmt.Fprintf(os.Stderr, "y0 y1 y2 y3 y4 = %v %v %v %v %v | %v %v %v %v %v\n", y0, y1, y2, y3, y4, _y0, _y1, _y2, _y3, _y4)

	return pack(y0, y1, y2, y3, y4)
}
