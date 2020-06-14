/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 *
 * https://leetcode.com/problems/merge-sorted-array/description/
 *
 * algorithms
 * Easy (38.88%)
 * Likes:    2090
 * Dislikes: 4045
 * Total Accepted:    573.2K
 * Total Submissions: 1.5M
 * Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
 *
 * Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as
 * one sorted array.
 * 
 * Note:
 * 
 * 
 * The number of elements initialized in nums1 and nums2 are m and n
 * respectively.
 * You may assume that nums1 has enough space (size that is greater or equal to
 * m + n) to hold additional elements from nums2.
 * 
 * 
 * Example:
 * 
 * 
 * Input:
 * nums1 = [1,2,3,0,0,0], m = 3
 * nums2 = [2,5,6],       n = 3
 * 
 * Output: [1,2,2,3,5,6]
 * 
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
	i,j:=m-1,n-1;
	for end:=len(nums1)-1;end>=0&&j>=0;end--{
		if i>=0&&nums1[i]>nums2[j]{
			nums1[end]=nums1[i]
			i--
		}else{
			nums1[end]=nums2[j]
			j--
		}
	}
}
// @lc code=end

