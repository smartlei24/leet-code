//Given two binary strings a and b, return their sum as a binary string.
//
//
// Example 1:
// Input: a = "11", b = "1"
//Output: "100"
//
// Example 2:
// Input: a = "1010", b = "1011"
//Output: "10101"
//
//
// Constraints:
//
//
// 1 <= a.length, b.length <= 10⁴
// a and b consist only of '0' or '1' characters.
// Each string does not contain leading zeros except for the zero itself.
//
//
// Related Topics 位运算 数学 字符串 模拟 👍 1093 👎 0

package add_binary

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func addBinary(a string, b string) string {

	lenA, lenB := len(a), len(b)
	hasPlus := 0
	lenR := lenA
	if lenB > lenA {
		lenR = lenB
	}
	lenR++

	r := make([]uint8, lenR)
	for i := 1; lenA-i >= 0 || lenB-i >= 0; i++ {
		indexA := lenA - i
		indexB := lenB - i
		cA, cB := 0, 0
		if indexA >= 0 {
			if a[indexA] == '1' {
				cA = 1
			}
		}

		if indexB >= 0 {
			if b[indexB] == '1' {
				cB = 1
			}
		}

		bitAddr := cA + cB + hasPlus
		if bitAddr > 1 {
			hasPlus = 1
		} else {
			hasPlus = 0
		}

		if bitAddr == 1 || bitAddr == 3 {
			r[lenR-i] = '1'
		} else {
			r[lenR-i] = '0'
		}

	}

	if hasPlus > 0 {
		r[0] = '1'
		return string(r)
	} else {
		return string(r[1:])
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestAddBinary(t *testing.T) {
	a, b := "11", "1"
	expected := "100"
	actual := addBinary(a, b)

	if actual != expected {
		t.Errorf("actual is %v", actual)
	}
}
