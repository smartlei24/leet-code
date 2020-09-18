package leetcode

import (
	"testing"
)

func myAtoi(str string) int {
	result := 0
	isNegative := false
	isOverflow := false
	start := 0
	border := 2147483647

	for i := 0; i < len(str); i++ {
		if str[i] == 32 {
			start++
		} else if str[i] == 45 {
			isNegative = true
			border++
			start = i + 1
			break
		} else if str[i] == 43 {
			start = i + 1
			break
		} else {
			break
		}
	}

	for i := start; i < len(str); i++ {
		r := str[i] - 48
		if r >= 0 && r <= 9 {
			result = result*10 + int(r)
			isOverflow = result > border
			if isOverflow {
				result = border
				break
			}
		} else {
			break
		}
	}

	if isNegative {
		result = result * -1
	}
	return int(result)
}

func TestMyAtoi(t *testing.T) {
	input := "4193 with words"
	expect := 4193
	actual := myAtoi(input)
	if expect != actual {
		t.Errorf("actual is %d", actual)
	}
}
