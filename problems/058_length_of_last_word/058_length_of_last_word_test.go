//Given a string s consisting of words and spaces, return the length of the
//last word in the string.
//
// A word is a maximal substring consisting of non-space characters only.
//
//
// Example 1:
//
//
//Input: s = "Hello World"
//Output: 5
//Explanation: The last word is "World" with length 5.
//
//
// Example 2:
//
//
//Input: s = "   fly me   to   the moon  "
//Output: 4
//Explanation: The last word is "moon" with length 4.
//
//
// Example 3:
//
//
//Input: s = "luffy is still joyboy"
//Output: 6
//Explanation: The last word is "joyboy" with length 6.
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 10⁴
// s consists of only English letters and spaces ' '.
// There will be at least one word in s.
//
//
// Related Topics 字符串 👍 628 👎 0

package length_of_last_word

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLastWord(s string) int {
	blank := uint8(' ')

	fistNotBlankIndex := len(s) - 1
	for ; fistNotBlankIndex >= 0; fistNotBlankIndex-- {
		if s[fistNotBlankIndex] != blank {
			break
		}
	}

	if fistNotBlankIndex < 0 {
		return 0
	}

	head := fistNotBlankIndex - 1
	for ; head >= 0; head-- {
		if s[head] == blank {
			break
		}
	}

	return fistNotBlankIndex - head
}

//leetcode submit region end(Prohibit modification and deletion)
