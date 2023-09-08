//Given an integer array nums, return all the triplets [nums[i], nums[j], nums[
//k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
//
// Notice that the solution set must not contain duplicate triplets.
//
//
// Example 1:
//
//
//Input: nums = [-1,0,1,2,-1,-4]
//Output: [[-1,-1,2],[-1,0,1]]
//Explanation:
//nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
//nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
//nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
//The distinct triplets are [-1,0,1] and [-1,-1,2].
//Notice that the order of the output and the order of the triplets does not
//matter.
//
//
// Example 2:
//
//
//Input: nums = [0,1,1]
//Output: []
//Explanation: The only possible triplet does not sum up to 0.
//
//
// Example 3:
//
//
//Input: nums = [0,0,0]
//Output: [[0,0,0]]
//Explanation: The only possible triplet sums up to 0.
//
//
//
// Constraints:
//
//
// 3 <= nums.length <= 3000
// -10⁵ <= nums[i] <= 10⁵
//
//
// Related Topics 数组 双指针 排序 👍 6318 👎 0

package _3sum

import (
	"reflect"
	"sort"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	results := make([][]int, 0)
	l := len(nums)
	if l < 3 {
		return results
	}

	sort.Ints(nums)
	for i := 0; i < l; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := l - 1
		for left < right {
			sum := nums[left] + nums[i] + nums[right]
			if sum > 0 {
				right--
			} else if sum == 0 {
				results = append(results, []int{nums[i], nums[left], nums[right]})

				curRight := nums[right]
				for ; nums[right] == curRight && left < right; right-- {
				}

				curLeft := nums[left]
				for ; nums[left] == curLeft && left < right; left++ {
				}
			} else {
				left++
			}
		}
	}

	return results
}

//leetcode submit region end(Prohibit modification and deletion)

func Test3Sum(t *testing.T) {
	input := []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}
	expected := [][]int{{-4, 0, 4}, {-4, 1, 3}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1}}
	actual := threeSum(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual is %v", actual)
	}
}
