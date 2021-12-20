package linked_list

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1367. 二叉树中的列表
// https://leetcode-cn.com/problems/linked-list-in-binary-tree/
func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}
	var dfs func(node *ListNode, tNode *TreeNode) bool
	dfs = func(node *ListNode, tNode *TreeNode) bool {
		// 链表已经全部匹配完，匹配成功
		if node == nil {
			return true
		}
		//   // 二叉树访问到了空节点，匹配失败
		if nil == tNode {
			return false
		}
		// 当前匹配的二叉树上节点的值与链表节点的值不相等，匹配失败
		if node.Val != tNode.Val {
			return false
		}
		return dfs(node.Next, tNode.Left) || dfs(node.Next, tNode.Right)
	}

	return dfs(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)

}
