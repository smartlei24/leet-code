package leetcode

import (
	"strings"
	"testing"
)

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	cycle := numRows*2 - 2
	var chars strings.Builder
	chars.Grow(len(s))
	l := 0
	for row := 0; row < numRows; row++ {
		for i := 0; ; i++ {
			k1 := row + cycle*i
			if k1 < len(s) {
				chars.WriteByte(s[k1])
				l++
			} else {
				break
			}

			p := cycle - row*2
			if p > 0 && p < cycle {
				k2 := k1 + p
				if k2 < len(s) {
					chars.WriteByte(s[k2])
					l++
				} else {
					break
				}
			}
		}
	}
	return chars.String()
}

func TestConvert(t *testing.T) {
	input := "PAYPALISHIRING"
	numRows := 4
	expect := "PINALSIGYAHRPI"
	actual := convert(input, numRows)
	if expect != actual {
		t.Errorf("actual is %s", actual)
	}
}

func BenchmarkConvert(b *testing.B) {
	input := "PAYPALISHIRING"
	numRows := 4
	for i := 0; i < b.N; i++ {
		convert(input, numRows)
	}
}
