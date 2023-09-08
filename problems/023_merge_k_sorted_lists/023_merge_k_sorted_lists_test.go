//You are given an array of k linked-lists lists, each linked-list is sorted in
//ascending order.
//
// Merge all the linked-lists into one sorted linked-list and return it.
//
//
// Example 1:
//
//
//Input: lists = [[1,4,5],[1,3,4],[2,6]]
//Output: [1,1,2,3,4,4,5,6]
//Explanation: The linked-lists are:
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//merging them into one sorted list:
//1->1->2->3->4->4->5->6
//
//
// Example 2:
//
//
//Input: lists = []
//Output: []
//
//
// Example 3:
//
//
//Input: lists = [[]]
//Output: []
//
//
//
// Constraints:
//
//
// k == lists.length
// 0 <= k <= 10⁴
// 0 <= lists[i].length <= 500
// -10⁴ <= lists[i][j] <= 10⁴
// lists[i] is sorted in ascending order.
// The sum of lists[i].length will not exceed 10⁴.
//
//
// Related Topics 链表 分治 堆（优先队列） 归并排序 👍 2624 👎 0

package merge_k_sorted_lists

import (
	"math"
	"reflect"
	"testing"
)

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
func mergeKLists(lists []*ListNode) *ListNode {
	findTopSuccess, topNode := findMinNode(lists)
	if !findTopSuccess {
		return nil
	}

	curNode := topNode
	for findSuccess, minNode := findMinNode(lists); findSuccess; findSuccess, minNode = findMinNode(lists) {
		curNode.Next = minNode
		curNode = minNode
	}

	return topNode
}

func findMinNode(lists []*ListNode) (bool, *ListNode) {
	minValue := math.MaxInt
	minValueIndex := -1
	for index, node := range lists {
		if node == nil {
			continue
		}

		if node.Val < minValue {
			minValue = node.Val
			minValueIndex = index
		}
	}
	if minValueIndex >= 0 {
		r := (lists)[minValueIndex]
		lists[minValueIndex] = lists[minValueIndex].Next
		return true, r
	}
	return false, nil
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMergeKList(t *testing.T) {
	testCases := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}
	expected := []int{1, 1, 2, 3, 4, 4, 5, 6}

	input := make([]*ListNode, 0)
	for _, arr := range testCases {
		input = append(input, arrayToLinkList(arr))
	}
	runResult := mergeKLists(input)
	actual := listListToArray(runResult)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual is %v", actual)
	}

}

func arrayToLinkList(array []int) *ListNode {
	if len(array) < 1 {
		return nil
	}
	topNode := ListNode{
		Val: array[0],
	}
	curNode := &topNode

	for i := 1; i < len(array); i++ {
		node := ListNode{
			Val: array[i],
		}
		curNode.Next = &node
		curNode = &node
	}

	return &topNode
}

func listListToArray(node *ListNode) []int {
	r := make([]int, 0)
	if node == nil {
		return r
	}

	for curNode := node; curNode != nil; curNode = curNode.Next {
		r = append(r, curNode.Val)
	}
	return r
}
