/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */

package leetcode

import (
	"strconv"
	"testing"
)

// @lc code=start
func isMatch(s string, p string) bool {
	var memo = make(map[string]bool)
	return dp(s, len(s)-1, p, len(p)-1, memo)
}

func dp(s string, si int, p string, pi int, memo map[string]bool) bool {
	if pi < 0 {
		return si < 0
	}

	if si < 0 {
		if pi%2 == 0 {
			return false
		}
		for ; pi > 0; pi = pi - 2 {
			if p[pi] != '*' {
				return false
			}
		}
		return true
	}

	key := strconv.Itoa(si) + "_" + strconv.Itoa(pi)
	if v, ok := memo[key]; ok {
		return v
	}

	r := false
	if p[pi] == s[si] || p[pi] == '.' {
		r = dp(s, si-1, p, pi-1, memo)
	}

	if p[pi] == '*' {
		notMatch := dp(s, si, p, pi-2, memo)
		matchAtLeastOnce := pi > 0 && (s[si] == p[pi-1] || p[pi-1] == '.') && dp(s, si-1, p, pi, memo)
		r = notMatch || matchAtLeastOnce
	}

	memo[key] = r
	return r
}

// @lc code=end

func TestIsMatch(t *testing.T) {
	s := "aab"
	p := "c*a**b"
	expect := false
	actual := isMatch(s, p)
	if expect != actual {
		t.Errorf("actual is %t", actual)
	}
}
