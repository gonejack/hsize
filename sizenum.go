package main

import "fmt"

var NaN = fmt.Errorf("NaN")
var units = [...]string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
var base1K, _ = NewSizeNum().From("1024")
var precisions = 2

type SizeN struct {
	is []int8 // integers
	ds []int8 // decimals
}

func (n *SizeN) From(str string) (*SizeN, error) {
	if n.is == nil {
		n.is = make([]int8, 0, len(str)+2)
	} else {
		n.is = n.is[:0]
	}
	n.ds = make([]int8, precisions)
	if len(str) == 0 {
		return nil, NaN
	}
	for _, r := range str {
		v := r - '0'
		if v < 0 || v > 9 {
			return nil, NaN
		} else {
			n.is = append(n.is, int8(v))
		}
	}
	return n, nil
}
func (n *SizeN) Cmp(s *SizeN) (r int) {
	var d int8
	defer func() {
		if d >= 1 {
			r = 1
		}
		if d <= -1 {
			r = -1
		}
	}()

	// integer comparison
	nis, sis := n.integers(), s.integers()
	switch {
	case len(nis) > len(sis):
		d = 1
		return
	case len(nis) < len(sis):
		d = -1
		return
	default:
		for i := range nis {
			d = nis[i] - sis[i]
			if d != 0 {
				return
			}
		}
	}

	// decimal comparison
	for i := 0; i < precisions; i++ {
		d = n.ds[i] - s.ds[i]
		if d != 0 {
			return
		}
	}

	return
}
func (n *SizeN) Lt(s *SizeN) bool {
	return n.Cmp(s) < 0
}
func (n *SizeN) Lte(s *SizeN) bool {
	return n.Cmp(s) <= 0
}
func (n *SizeN) Gt(s *SizeN) bool {
	return n.Cmp(s) > 0
}
func (n *SizeN) Gte(s *SizeN) bool {
	return n.Cmp(s) >= 0
}
func (n *SizeN) Eq(s *SizeN) bool {
	return n.Cmp(s) == 0
}
func (n *SizeN) Div1024() {
	for i := 1; i <= 10; i++ {
		n.Div2()
	}
}
func (n *SizeN) Div2() {
	h := false
	for i, v := range n.is {
		if h {
			v += 10
		}
		if v > 0 {
			n.is[i], h = v>>1, v&1 != 0
		}
	}
	for i, v := range n.ds {
		if h {
			v += 10
		}
		if v > 0 {
			n.ds[i], h = v>>1, v&1 != 0
		}
	}
}
func (n *SizeN) String() string {
	arr := make([]rune, 0, len(n.is)+precisions+1)
	for _, v := range n.integers() {
		arr = append(arr, rune('0'+v))
	}
	if len(arr) == 0 {
		arr = append(arr, '0')
	}
	for i, v := range n.decimals() {
		if i == 0 {
			arr = append(arr, '.')
		}
		arr = append(arr, rune('0'+v))
	}
	return string(arr)
}

func (n *SizeN) integers() (is []int8) {
	for i, v := range n.is {
		if v != 0 {
			return n.is[i:]
		}
	}
	return
}
func (n *SizeN) decimals() (ds []int8) {
	for i := precisions; i > 0; i-- {
		if n.ds[i-1] != 0 {
			return n.ds[:i]
		}
	}
	return
}

func NewSizeNum() *SizeN {
	return new(SizeN)
}
