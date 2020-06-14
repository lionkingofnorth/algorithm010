/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 *
 * https://leetcode.com/problems/plus-one/description/
 *
 * algorithms
 * Easy (42.13%)
 * Likes:    1436
 * Dislikes: 2291
 * Total Accepted:    583.4K
 * Total Submissions: 1.4M
 * Testcase Example:  '[1,2,3]'
 *
 * Given a non-empty array of digits representing a non-negative integer, plus
 * one to the integer.
 * 
 * The digits are stored such that the most significant digit is at the head of
 * the list, and each element in the array contain a single digit.
 * 
 * You may assume the integer does not contain any leading zero, except the
 * number 0 itself.
 * 
 * Example 1:
 * 
 * 
 * Input: [1,2,3]
 * Output: [1,2,4]
 * Explanation: The array represents the integer 123.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [4,3,2,1]
 * Output: [4,3,2,2]
 * Explanation: The array represents the integer 4321.
 * 
 */

// @lc code=start
func plusOne(digits []int) []int {
    for end:=len(digits)-1;end>=0;end--{
		if digits[end]!=9{
			digits[end]+=1
			break
		}else{
			digits[end]=0
		}
	}
	if digits[0]==0{digits=append([]int{1},digits[:]...)}
	return digits
}
// @lc code=end

