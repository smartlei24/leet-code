//Given a sorted array of distinct integers and a target value, return the
//index if the target is found. If not, return the index where it would be if it were
//inserted in order.
//
// You must write an algorithm with O(log n) runtime complexity.
//
//
// Example 1:
//
//
//Input: nums = [1,3,5,6], target = 5
//Output: 2
//
//
// Example 2:
//
//
//Input: nums = [1,3,5,6], target = 2
//Output: 1
//
//
// Example 3:
//
//
//Input: nums = [1,3,5,6], target = 7
//Output: 4
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
// nums contains distinct values sorted in ascending order.
// -10⁴ <= target <= 10⁴
//
//
// Related Topics 数组 二分查找 👍 2140 👎 0

package search_insert_position

// leetcode submit region begin(Prohibit modification and deletion)
func searchInsert(nums []int, target int) int {
	lens := len(nums)

	left, right := 0, lens-1

	for right >= left {
		compareIndex := (left + right) / 2
		compareNum := nums[compareIndex]
		if compareNum == target {
			return compareIndex
		} else if compareNum > target {
			right = compareIndex - 1
		} else {
			left = compareIndex + 1
		}
	}
	return left
}

//leetcode submit region end(Prohibit modification and deletion)
