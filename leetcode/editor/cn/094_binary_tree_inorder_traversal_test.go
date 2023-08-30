//Given the root of a binary tree, return the inorder traversal of its nodes'
//values.
//
//
// Example 1:
//
//
//Input: root = [1,null,2,3]
//Output: [1,3,2]
//
//
// Example 2:
//
//
//Input: root = []
//Output: []
//
//
// Example 3:
//
//
//Input: root = [1]
//Output: [1]
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
//
//Follow up: Recursive solution is trivial, could you do it iteratively?
//
// Related Topics 栈 树 深度优先搜索 二叉树 👍 1897 👎 0

package problems

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	slice := make([]int, 0)

	inOrder(root, &slice)
	return slice
}

func inOrder(node *TreeNode, slice *[]int) {
	if node == nil {
		return
	}

	if node.Left != nil {
		inOrder(node.Left, slice)
	}

	*slice = append(*slice, node.Val)

	if node.Right != nil {
		inOrder(node.Right, slice)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
