/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 *
 * https://leetcode.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (55.72%)
 * Likes:    3425
 * Dislikes: 180
 * Total Accepted:    671.7K
 * Total Submissions: 1.2M
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * Given an array of strings, group anagrams together.
 * 
 * Example:
 * 
 * 
 * Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
 * Output:
 * [
 * ⁠ ["ate","eat","tea"],
 * ⁠ ["nat","tan"],
 * ⁠ ["bat"]
 * ]
 * 
 * Note:
 * 
 * 
 * All inputs will be in lowercase.
 * The order of your output does not matter.
 * 
 * 
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	sumMap:=make(map[string]int)
	var rowCnt int=0
	var res [][]string
	for _,v:=range strs{
		cur:=orderString(v)
		if row,ok:=sumMap[cur];ok{
			res[row]=append(res[row],v)
			continue
		}else{
			sumMap[cur]=rowCnt
			if len(res)!=rowCnt+1{
				res=append(res,[]string{})
			}
			res[rowCnt]=append(res[rowCnt],v)
			rowCnt++
		}
	}
	return res	
}

func orderString(s string)string{

rs:=[]rune(s)

sort.Slice(rs,func(a,b int)bool{return rs[a]<rs[b]})
return string(rs)
}
// @lc code=end

