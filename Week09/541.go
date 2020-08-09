/*
 * @lc app=leetcode id=541 lang=golang
 *
 * [541] Reverse String II
 *
 * https://leetcode.com/problems/reverse-string-ii/description/
 *
 * algorithms
 * Easy (47.97%)
 * Likes:    443
 * Dislikes: 1259
 * Total Accepted:    90.7K
 * Total Submissions: 187.3K
 * Testcase Example:  '"abcdefg"\n2'
 *
 * 
 * Given a string and an integer k, you need to reverse the first k characters
 * for every 2k characters counting from the start of the string. If there are
 * less than k characters left, reverse all of them. If there are less than 2k
 * but greater than or equal to k characters, then reverse the first k
 * characters and left the other as original.
 * 
 * 
 * Example:
 * 
 * Input: s = "abcdefg", k = 2
 * Output: "bacdfeg"
 * 
 * 
 * 
 * Restrictions: 
 * 
 * ⁠The string consists of lower English letters only.
 * ⁠Length of the given string and k will in the range [1, 10000]
 * 
 */

// @lc code=start
func reverseStr(s string, k int) string {
	sb := []byte(s) // in Go string is immutable. Need to convert s to byte array
	slen := len(sb)

	shouldReverse := true
	for i := 0; i < slen; i = i + k {
		if shouldReverse {
			end := i + k - 1
			if end >= slen-1 {
				end = slen - 1
			}
			reverseSb(sb, i, end)
		}
		shouldReverse = !shouldReverse
	}
	return string(sb)
}

func reverseSb(sb []byte, start, end int) {
	for start < end {
		sb[start], sb[end] = sb[end], sb[start]
		start++
		end--
	}
}
// @lc code=end

