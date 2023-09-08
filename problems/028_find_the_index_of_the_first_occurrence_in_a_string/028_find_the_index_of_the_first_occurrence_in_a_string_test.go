//Given two strings needle and haystack, return the index of the first
//occurrence of needle in haystack, or -1 if needle is not part of haystack.
//
//
// Example 1:
//
//
//Input: haystack = "sadbutsad", needle = "sad"
//Output: 0
//Explanation: "sad" occurs at index 0 and 6.
//The first occurrence is at index 0, so we return 0.
//
//
// Example 2:
//
//
//Input: haystack = "leetcode", needle = "leeto"
//Output: -1
//Explanation: "leeto" did not occur in "leetcode", so we return -1.
//
//
//
// Constraints:
//
//
// 1 <= haystack.length, needle.length <= 10⁴
// haystack and needle consist of only lowercase English characters.
//
//
// Related Topics 双指针 字符串 字符串匹配 👍 1984 👎 0

package find_the_index_of_the_first_occurrence_in_a_string

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func strStr(haystack string, needle string) int {

	lengthH := len(haystack)
	lengthN := len(needle)
	if lengthH < lengthN {
		return -1
	}
	matchWidth := 0
	for left := 0; left < lengthH; left++ {
		for matchWidth = 0; left+matchWidth < lengthH && matchWidth < lengthN; matchWidth++ {
			compareIndex := left + matchWidth
			if haystack[compareIndex] != needle[matchWidth] {
				break
			}
		}
		if matchWidth == lengthN {
			return left
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestStrStr(t *testing.T) {
	input := "sadbutsad"
	search := "sad"
	expected := 0
	actual := strStr(input, search)
	if actual != expected {
		t.Errorf("actual is %v", actual)
	}
}
