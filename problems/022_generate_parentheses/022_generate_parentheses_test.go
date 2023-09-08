//Given n pairs of parentheses, write a function to generate all combinations
//of well-formed parentheses.
//
//
// Example 1:
// Input: n = 3
//Output: ["((()))","(()())","(())()","()(())","()()()"]
//
// Example 2:
// Input: n = 1
//Output: ["()"]
//
//
// Constraints:
//
//
// 1 <= n <= 8
//
//
// Related Topics 字符串 动态规划 回溯 👍 3367 👎 0

package generate_parentheses

// leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	result := make([]string, 0, 0)
	addParenthesis(&result, "", 0, 0, n)
	return result
}

func addParenthesis(result *[]string, s string, left int, right int, n int) {
	if left > n || right > n {
		return
	}

	if left == n && right == n {
		*result = append(*result, s)
	}

	if left < n {
		v := s + "("
		addParenthesis(result, v, left+1, right, n)
	}

	if right < left {
		v := s + ")"
		addParenthesis(result, v, left, right+1, n)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
