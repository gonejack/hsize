package main

import (
	"fmt"
)

var prec = 2
var NaN = fmt.Errorf("NaN")

type SizeNum struct {
	ints []int8
	decs []int8
}

func (n *SizeNum) From(str string) (*SizeNum, error) {
	if n.ints == nil {
		n.ints = make([]int8, 0, len(str)+2)
	} else {
		n.ints = n.ints[0:0]
	}
	n.decs = make([]int8, prec, prec)

	if len(str) == 0 {
		return nil, NaN
	}

	for _, r := range str {
		v := r - '0'
		if v < 0 || v > 9 {
			return nil, NaN
		} else {
			n.ints = append(n.ints, int8(v))
		}
	}

	return n, nil
}

func (n *SizeNum) Cmp(s *SizeNum) (result int) {
	var d int8
	defer func() {
		if d >= 1 {
			result = 1
		}
		if d <= -1 {
			result = -1
		}
	}()

	// integer comparison
	nints, sints := n.integers(), s.integers()
	switch {
	case len(nints) > len(sints):
		d = 1
		return
	case len(nints) < len(sints):
		d = -1
		return
	default:
		for i := range nints {
			d = nints[i] - sints[i]
			if d != 0 {
				return
			}
		}
	}

	// decimal comparison
	for i := 0; i < prec; i++ {
		d = n.decs[i] - s.decs[i]
		if d != 0 {
			return
		}
	}

	return
}
func (n *SizeNum) Lt(s *SizeNum) bool {
	return n.Cmp(s) < 0
}
func (n *SizeNum) Lte(s *SizeNum) bool {
	return n.Cmp(s) <= 0
}
func (n *SizeNum) Gt(s *SizeNum) bool {
	return n.Cmp(s) > 0
}
func (n *SizeNum) Gte(s *SizeNum) bool {
	return n.Cmp(s) >= 0
}
func (n *SizeNum) Eq(s *SizeNum) bool {
	return n.Cmp(s) == 0
}

func (n *SizeNum) Div1024() {
	for i := 1; i <= 10; i++ {
		n.Div2()
	}
}
func (n *SizeNum) Div2() {
	h := false
	for i, v := range n.ints {
		if h {
			v += 10
		}
		if v > 0 {
			n.ints[i], h = v>>1, v&1 != 0
		}
	}
	for i, v := range n.decs {
		if h {
			v += 10
		}
		if v > 0 {
			n.decs[i], h = v>>1, v&1 != 0
		}
	}
}
func (n *SizeNum) String() string {
	rs := make([]rune, 0, len(n.ints)+prec+1)
	ints, decs := n.integers(), n.decimals()

	if len(ints) == 0 {
		ints = append(ints, '0')
	} else {
		for _, v := range ints {
			rs = append(rs, rune('0'+v))
		}
	}
	if len(decs) > 0 {
		rs = append(rs, '.')
	}
	for _, v := range decs {
		rs = append(rs, rune('0'+v))
	}
	return string(rs)
}

func (n *SizeNum) integers() (ints []int8) {
	for i, v := range n.ints {
		if v != 0 {
			return n.ints[i:]
		}
	}
	return
}
func (n *SizeNum) decimals() (decs []int8) {
	for i := prec; i > 0; i-- {
		if n.decs[i-1] > 0 {
			return n.decs[:i]
		}
	}
	return
}
