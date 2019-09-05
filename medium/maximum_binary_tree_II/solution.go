package maximum_binary_tree_II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	// safety catch
	if root == nil {
		return &TreeNode{Val: val}
	}

	// we have a new root
	if val > root.Val {
		return &TreeNode{
			Val:  val,
			Left: root,
		}
	}

	// if both are nil, then the root
	// in A was a leaf node (and possibly lone)
	if root.Right == nil && root.Left == nil {
		root.Right = &TreeNode{Val: val}
		return root
	}

	root.Right = insertIntoMaxTree(root.Right, val)
	return root
}
