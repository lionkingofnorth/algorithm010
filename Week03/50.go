/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 *
 * https://leetcode.com/problems/powx-n/description/
 *
 * algorithms
 * Medium (29.46%)
 * Likes:    1456
 * Dislikes: 2968
 * Total Accepted:    461.8K
 * Total Submissions: 1.6M
 * Testcase Example:  '2.00000\n10'
 *
 * Implement pow(x, n), which calculates x raised to the power n (x^n).
 * 
 * Example 1:
 * 
 * 
 * Input: 2.00000, 10
 * Output: 1024.00000
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 2.10000, 3
 * Output: 9.26100
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: 2.00000, -2
 * Output: 0.25000
 * Explanation: 2^-2 = 1/2^2 = 1/4 = 0.25
 * 
 * 
 * Note:
 * 
 * 
 * -100.0 < x < 100.0
 * n is a 32-bit signed integer, within the range [−2^31, 2^31 − 1]
 * 
 * 
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	
	//terminator
	if x==0{
		return 0
	}
	if n==0{
		return 1
	}

	//first check 
	xstar,nstar:=mincheck(x,n)

	//half
	half:=myPow(xstar,nstar/2)

	if nstar%2!=0{
		return half*half*xstar
	}else{	
		return half*half
	}
}

func mincheck(x float64,n int) (float64,int){
	if n<0{return 1/x,-n}else{
		return x,n
	}
}
// @lc code=end

