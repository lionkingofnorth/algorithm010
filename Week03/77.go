/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 *
 * https://leetcode.com/problems/combinations/description/
 *
 * algorithms
 * Medium (53.50%)
 * Likes:    1418
 * Dislikes: 66
 * Total Accepted:    285.6K
 * Total Submissions: 529.2K
 * Testcase Example:  '4\n2'
 *
 * Given two integers n and k, return all possible combinations of k numbers
 * out of 1 ... n.
 * 
 * Example:
 * 
 * 
 * Input: n = 4, k = 2
 * Output:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 * 
 * 
 */

// @lc code=start
func combine(n int, k int) [][]int {
    results := [][]int{}
    if k > n {
        return results
    }
    
    dfs([]int{}, n, k, 1, &results)
    return results
}

func dfs(buf []int, n, k, idx int, results *[][]int) {
    if k == 0 {
        tmp := make([]int, len(buf))
        copy(tmp, buf)
        *results = append(*results, tmp)
    }
    
    for i := idx; i <= n; i++ {
        buf = append(buf, i)
        
        dfs(buf, n, k-1, i+1, results)
        
        buf = buf[:len(buf)-1]
    }
}
// @lc code=end

