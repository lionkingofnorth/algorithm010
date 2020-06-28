/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
 *
 * https://leetcode.com/problems/binary-tree-postorder-traversal/description/
 *
 * algorithms
 * Hard (53.84%)
 * Likes:    1682
 * Dislikes: 88
 * Total Accepted:    373.3K
 * Total Submissions: 688.3K
 * Testcase Example:  '[1,null,2,3]'
 *
 * Given a binary tree, return the postorder traversal of its nodes' values.
 * 
 * Example:
 * 
 * 
 * Input: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 * 
 * Output: [3,2,1]
 * 
 * 
 * Follow up: Recursive solution is trivial, could you do it iteratively?
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func dfs(node *TreeNode,res []int)[]int{
	if node==nil{
		return res
	}
	var left,right []int
	//left binary tree
	if node.Left!=nil{
		left=dfs(node.Left,res)
	}
	//right binary tree
	if node.Right!=nil{
		right=dfs(node.Right,res)
	}

	//inorder search
	res=append(res,left...)
	res=append(res,right...)
	res=append(res,node.Val)

	//return 
	return res
}
func postorderTraversal(root *TreeNode) []int {
    return dfs(root,[]int{})
}
// @lc code=end

