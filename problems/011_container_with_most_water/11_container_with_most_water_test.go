package leetcode

import "testing"

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
	max := 0

	for left, right := 0, len(height)-1; left < right; {
		var h int
		w := right - left
		if height[left] < height[right] {
			h = height[left]
			left++
		} else {
			h = height[right]
			right--
		}

		area := w * h
		if area > max {
			max = area
		}
	}

	return max
}

// @lc code=end

func TestMaxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	expect := 49

	actual := maxArea(height)
	if actual != expect {
		t.Errorf("actual is %d", actual)
	}
}
