//You are given a large integer represented as an integer array digits, where
//each digits[i] is the iᵗʰ digit of the integer. The digits are ordered from most
//significant to least significant in left-to-right order. The large integer does
//not contain any leading 0's.
//
// Increment the large integer by one and return the resulting array of digits.
//
//
//
// Example 1:
//
//
//Input: digits = [1,2,3]
//Output: [1,2,4]
//Explanation: The array represents the integer 123.
//Incrementing by one gives 123 + 1 = 124.
//Thus, the result should be [1,2,4].
//
//
// Example 2:
//
//
//Input: digits = [4,3,2,1]
//Output: [4,3,2,2]
//Explanation: The array represents the integer 4321.
//Incrementing by one gives 4321 + 1 = 4322.
//Thus, the result should be [4,3,2,2].
//
//
// Example 3:
//
//
//Input: digits = [9]
//Output: [1,0]
//Explanation: The array represents the integer 9.
//Incrementing by one gives 9 + 1 = 10.
//Thus, the result should be [1,0].
//
//
//
// Constraints:
//
//
// 1 <= digits.length <= 100
// 0 <= digits[i] <= 9
// digits does not contain any leading 0's.
//
//
// Related Topics 数组 数学 👍 1285 👎 0

package plus_one

import (
	"reflect"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func plusOne(digits []int) []int {
	p := false
	l := len(digits)

	for i := l - 1; i >= 0; i-- {
		addValue := 0
		if p {
			addValue++
		}
		if i == l-1 {
			addValue++
		}
		if addValue == 0 {
			return digits
		}
		v := digits[i] + addValue

		p = v >= 10
		if p {
			v = v - 10
		}
		digits[i] = v
	}

	if p {
		result := make([]int, l+1)
		result[0] = 1
		return result
	}

	return digits

}

//leetcode submit region end(Prohibit modification and deletion)

func TestPlusOne(t *testing.T) {
	testCase := []int{9, 9}
	expected := []int{1, 0, 0}
	actual := plusOne(testCase)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual is %v", actual)
	}
}
