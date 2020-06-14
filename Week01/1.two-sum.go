/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 *
 * https://leetcode.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (45.40%)
 * Likes:    14943
 * Dislikes: 540
 * Total Accepted:    2.9M
 * Total Submissions: 6.4M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers, return indices of the two numbers such that they
 * add up to a specific target.
 * 
 * You may assume that each input would have exactly one solution, and you may
 * not use the same element twice.
 * 
 * Example:
 * 
 * 
 * Given nums = [2, 7, 11, 15], target = 9,
 * 
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 * 
 * 
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	lookup := make(map[int]int)
    for i, v := range nums {
		fmt.Println("-v:",-v)
		j, ok := lookup[-v]
		fmt.Println("lookup:",lookup[-v],ok)
        lookup[v - target] = i
        if ok {
            return []int{j, i}
        }
    }
    return []int{}
}
// @lc code=end

