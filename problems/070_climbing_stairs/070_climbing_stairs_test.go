//You are climbing a staircase. It takes n steps to reach the top.
//
// Each time you can either climb 1 or 2 steps. In how many distinct ways can
//you climb to the top?
//
//
// Example 1:
//
//
//Input: n = 2
//Output: 2
//Explanation: There are two ways to climb to the top.
//1. 1 step + 1 step
//2. 2 steps
//
//
// Example 2:
//
//
//Input: n = 3
//Output: 3
//Explanation: There are three ways to climb to the top.
//1. 1 step + 1 step + 1 step
//2. 1 step + 2 steps
//3. 2 steps + 1 step
//
//
//
// Constraints:
//
//
// 1 <= n <= 45
//
//
// Related Topics 记忆化搜索 数学 动态规划 👍 3227 👎 0

package climbing_stairs

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	n1, n2 := 1, 2
	for i := 3; i < n; i++ {
		n1, n2 = n2, n1+n2
	}
	return n1 + n2
}

//leetcode submit region end(Prohibit modification and deletion)

func TestClimbStairs(t *testing.T) {
	n := 3
	expected := 3
	actual := climbStairs(n)
	if expected != actual {
		t.Errorf("actual is %v", actual)
	}
}
