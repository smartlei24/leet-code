//Given the root of a binary tree, return the postorder traversal of its nodes'
//values.
//
//
// Example 1:
//
//
//Input: root = [1,null,2,3]
//Output: [3,2,1]
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
// The number of the nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100
//
//
//
//Follow up: Recursive solution is trivial, could you do it iteratively?
//
// Related Topics 栈 树 深度优先搜索 二叉树 👍 1080 👎 0

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
func postorderTraversal(root *TreeNode) []int {
	slice := make([]int, 0)

	postOrder(root, &slice)
	return slice
}

func postOrder(node *TreeNode, slice *[]int) {
	if node == nil {
		return
	}

	if node.Left != nil {
		postOrder(node.Left, slice)
	}

	if node.Right != nil {
		postOrder(node.Right, slice)
	}

	*slice = append(*slice, node.Val)
}

//leetcode submit region end(Prohibit modification and deletion)
