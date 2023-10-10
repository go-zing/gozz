package zutils

import (
	"testing"
)

func TestSplitKV(t *testing.T) {
	for _, c := range [][4]string{
		{"k=v=3", "=", "k", "v=3"},
		{"k=v", "=", "k", "v"},
		{"k", "=", "k", ""},
		{"", "=", "", ""},
	} {
		k, v := SplitKV(c[0], c[1])
		if k != c[2] || v != c[3] {
			t.Fatalf("want (%v,%v) got (%v,%v)", c[2], c[3], k, v)
		}
	}
}
