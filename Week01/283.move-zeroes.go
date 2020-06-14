/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int)  {
	j:=0
    for i:=range nums{
		if nums[i]!=0{
			nums[j],nums[i]=nums[i],nums[j]
			j++
		}
	}

}
// @lc code=end

