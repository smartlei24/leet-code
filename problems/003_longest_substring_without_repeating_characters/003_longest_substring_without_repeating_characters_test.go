package problems

import "testing"

func lengthOfLongestSubstring(s string) int {

	var max, curLength, start int

	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		p := s[i]
		if old, ok := m[p]; !ok || old < start {
			m[p] = i
			curLength++
		} else {
			if curLength > max {
				max = curLength
			}
			start = old
			curLength = i - old
			m[p] = i
		}
	}

	if curLength > max {
		max = curLength
	}
	return max
}

func lengthOfLongestSubstring2(s string) int {
	dict := make([]bool, 256)
	left, right := 0, 0
	res := 0

	for right < len(s) {
		for left < right && dict[s[right]] {
			dict[s[left]] = false
			left++
		}

		dict[s[right]] = true
		res = maxInt(res, right-left+1)
		right++
	}
	return res
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "pwwkew"
	r := lengthOfLongestSubstring2(s)
	t.Logf("actual is %d", r)
}
