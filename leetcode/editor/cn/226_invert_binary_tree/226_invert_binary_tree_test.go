//Given the root of a binary tree, invert the tree, and return its root.
//
//
// Example 1:
//
//
//Input: root = [4,2,7,1,3,6,9]
//Output: [4,7,2,9,6,3,1]
//
//
// Example 2:
//
//
//Input: root = [2,1,3]
//Output: [2,3,1]
//
//
// Example 3:
//
//
//Input: root = []
//Output: []
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1662 👎 0

package invert_binary_tree

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
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invertTree(root.Right)
	invertTree(root.Left)
	right := root.Right
	root.Right = root.Left
	root.Left = right
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
