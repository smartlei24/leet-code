//Given the head of a sorted linked list, delete all duplicates such that each
//element appears only once. Return the linked list sorted as well.
//
//
// Example 1:
//
//
//Input: head = [1,1,2]
//Output: [1,2]
//
//
// Example 2:
//
//
//Input: head = [1,1,2,3,3]
//Output: [1,2,3]
//
//
//
// Constraints:
//
//
// The number of nodes in the list is in the range [0, 300].
// -100 <= Node.val <= 100
// The list is guaranteed to be sorted in ascending order.
//
//
// Related Topics 链表 👍 1036 👎 0

package remove_duplicates_from_sorted_list

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	cur := head
	for cur.Next != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		}

		if cur.Next != nil {
			cur = cur.Next
		}
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
