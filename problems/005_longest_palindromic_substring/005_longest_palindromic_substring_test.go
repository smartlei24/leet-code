package problems

import (
	"testing"
)

func longestPalindrome(s string) string {
	maxStr := ""
	for i := 0; i < len(s); i++ {
		s1 := expandAroundIndex(s, i, i)
		s2 := expandAroundIndex(s, i, i+1)
		if len(s1) > len(maxStr) {
			maxStr = s1
		}
		if len(s2) > len(maxStr) {
			maxStr = s2
		}
	}
	return maxStr
}

func expandAroundIndex(s string, i int, j int) string {
	l := 0
	for i >= 0 && j < len(s) {
		if s[i] == s[j] {
			l = j - i + 1
			i--
			j++
			continue
		}
		break
	}
	return s[i+1 : i+1+l]
}

func TestLongestPalindrome(t *testing.T) {
	s := "babad"
	expect := "bab"
	actual := longestPalindrome(s)
	if expect != actual {
		t.Errorf("actual is %s", actual)
	}
}
