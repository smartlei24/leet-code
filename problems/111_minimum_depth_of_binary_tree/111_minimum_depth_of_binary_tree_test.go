//Given a binary tree, find its minimum depth.
//
// The minimum depth is the number of nodes along the shortest path from the
//root node down to the nearest leaf node.
//
// Note: A leaf is a node with no children.
//
//
// Example 1:
//
//
//Input: root = [3,9,20,null,null,15,7]
//Output: 2
//
//
// Example 2:
//
//
//Input: root = [2,null,3,null,4,null,5,null,6]
//Output: 5
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [0, 10⁵].
// -1000 <= Node.val <= 1000
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1070 👎 0

package minimum_depth_of_binary_tree

import "math"

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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Right == nil && root.Left == nil {
		return 1
	}

	leftMinDepth, rightMinDepth := math.MaxInt, math.MaxInt
	if root.Left != nil {
		leftMinDepth = minDepth(root.Left)
	}

	if root.Right != nil {
		rightMinDepth = minDepth(root.Right)
	}
	return min(leftMinDepth, rightMinDepth) + 1
}

func min(m1 int, m2 int) int {
	if m1 > m2 {
		return m2
	}
	return m1
}

//leetcode submit region end(Prohibit modification and deletion)
