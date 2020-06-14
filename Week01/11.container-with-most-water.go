/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start

func min(a,b int)int{
	if a<b{return a}
	return b
}
func max(a,b int)int{
	if a>b{return a}
	return b
}


func maxArea(height []int) int {
	i,j:=0,len(height)-1
	var area int
	for i<j{
		area=max(area,(j-i)*min(height[i],height[j]))
		if height[i]<height[j]{
			i++
		}else{
			j--
		}
	}
	return area
}
// @lc code=end

