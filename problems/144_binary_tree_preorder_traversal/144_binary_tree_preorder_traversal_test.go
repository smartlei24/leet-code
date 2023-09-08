//Given the root of a binary tree, return the preorder traversal of its nodes'
//values.
//
//
// Example 1:
//
//
//Input: root = [1,null,2,3]
//Output: [1,2,3]
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
// Follow up: Recursive solution is trivial, could you do it iteratively?
//
// Related Topics 栈 树 深度优先搜索 二叉树 👍 1120 👎 0

package binary_tree_preorder_traversal_test

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)

func preorderTraversal(root *TreeNode) []int {
	slice := make([]int, 0)

	preorder(root, &slice)
	return slice
}

func preorder(node *TreeNode, slice *[]int) {
	if node == nil {
		return
	}

	*slice = append(*slice, node.Val)
	if node.Left != nil {
		preorder(node.Left, slice)
	}

	if node.Right != nil {
		preorder(node.Right, slice)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
