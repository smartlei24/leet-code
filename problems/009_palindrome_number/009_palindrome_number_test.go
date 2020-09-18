package leetcode

import "testing"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reverse := 0
	n := x

	for {
		if n < 10 {
			reverse = reverse*10 + n
			break
		}
		temp := n / 10
		p := n - temp*10
		n = temp
		reverse = reverse*10 + p
	}
	return reverse == x
}

func TestIsPalindrome(t *testing.T) {
	inputs := []int{121, -121, 10}
	expects := []bool{true, false, false}
	for i := 0; i < len(inputs); i++ {
		actual := isPalindrome(inputs[i])
		if actual != expects[i] {
			t.Errorf("input is %d, expect is %t, but actual is %t", inputs[i], expects[i], actual)
		}
	}
}
