package main

import "testing"

func TestMinimum(t *testing.T) {
	m := NewModel(NewConfig())
	in := []int64{5, 4, 3, 2, 1}
	exIdx := 4

	min, idx := m.getMinimumIn(in)

	equals(t, idx, 4)
	equals(t, min, in[exIdx])
}
