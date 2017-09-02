package sort

import (
	"runtime"
	"testing"
)

func TestParallel(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{10, 11, 13, 2, 20, 3, 0, 40, 1, 30, 23, 2, 1}, []int{0, 1, 1, 2, 2, 3, 10, 11, 13, 20, 23, 30, 40}},
	}

	for _, c := range cases {
		got := Parallel(c.in, runtime.NumCPU())
		if !Equals(got, c.want) {
			t.Errorf("Parallel(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
