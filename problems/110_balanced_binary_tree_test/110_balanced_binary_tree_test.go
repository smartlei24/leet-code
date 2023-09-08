//Given a binary tree, determine if it is height-balanced.
//
//
// Example 1:
//
//
//Input: root = [3,9,20,null,null,15,7]
//Output: true
//
//
// Example 2:
//
//
//Input: root = [1,2,2,3,3,null,null,4,4]
//Output: false
//
//
// Example 3:
//
//
//Input: root = []
//Output: true
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [0, 5000].
// -10⁴ <= Node.val <= 10⁴
//
//
// Related Topics 树 深度优先搜索 二叉树 👍 1392 👎 0

package balanced_binary_tree_test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !isBalanced(root.Left) {
		return false
	}
	if !isBalanced(root.Right) {
		return false
	}

	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)

	var heightDiff = leftHeight - rightHeight
	return heightDiff < 2 && heightDiff > -2

}

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight, rightHeight := 0, 0
	if node.Left != nil {
		leftHeight = getHeight(node.Left)
	}
	if node.Right != nil {
		rightHeight = getHeight(node.Right)
	}
	return max(rightHeight, leftHeight) + 1
}

func max(m1 int, m2 int) int {
	if m1 < m2 {
		return m2
	}
	return m1
}

//leetcode submit region end(Prohibit modification and deletion)
