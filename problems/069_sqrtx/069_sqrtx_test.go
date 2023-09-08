//Given a non-negative integer x, return the square root of x rounded down to
//the nearest integer. The returned integer should be non-negative as well.
//
// You must not use any built-in exponent function or operator.
//
//
// For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.
//
//
//
// Example 1:
//
//
//Input: x = 4
//Output: 2
//Explanation: The square root of 4 is 2, so we return 2.
//
//
// Example 2:
//
//
//Input: x = 8
//Output: 2
//Explanation: The square root of 8 is 2.82842..., and since we round it down
//to the nearest integer, 2 is returned.
//
//
//
// Constraints:
//
//
// 0 <= x <= 2³¹ - 1
//
//
// Related Topics 数学 二分查找 👍 1428 👎 0

package sqrtx

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	if x == 1 {
		return 1
	}

	left, right := 1, x/2
	for left <= right {
		mid := (left + right) / 2
		pow := mid * mid

		if pow == x {
			return mid
		} else if pow < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left - 1

}

//leetcode submit region end(Prohibit modification and deletion)

func TestMySqrt(t *testing.T) {
	x := 169
	expected := 13
	actual := mySqrt(x)
	if actual != expected {
		t.Errorf("actual is %v", actual)
	}
}
