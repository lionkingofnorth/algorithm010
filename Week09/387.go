/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 *
 * https://leetcode.com/problems/first-unique-character-in-a-string/description/
 *
 * algorithms
 * Easy (53.09%)
 * Likes:    2021
 * Dislikes: 117
 * Total Accepted:    558.4K
 * Total Submissions: 1M
 * Testcase Example:  '"leetcode"'
 *
 * Given a string, find the first non-repeating character in it and return its
 * index. If it doesn't exist, return -1.
 * 
 * Examples:
 * 
 * 
 * s = "leetcode"
 * return 0.
 * 
 * s = "loveleetcode"
 * return 2.
 * 
 * 
 * 
 * 
 * Note: You may assume the string contains only lowercase English letters.
 * 
 */

// @lc code=start
func firstUniqChar(s string) int {
	magic := [26]int{}
	for i := range magic {
		magic[i] = -1
	}
	for i, v := range s {
		if magic[v-'a'] == -1 {
			magic[v-'a'] = i
		} else {
			magic[v-'a'] = -2
		}
	}
	ret := len(s) + 1
	for _, v := range magic {
		if v >= 0 {
			if v < ret {
				ret = v
			}
		}
	}
	if ret == len(s) + 1 {
		ret = -1
	}
	return ret
}
// @lc code=end

