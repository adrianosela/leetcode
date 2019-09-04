package lowest_common_ancestor_of_a_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p {
		if hasChildren(root, q) {
			return root
		}
	}

	if root == q {
		if hasChildren(root, p) {
			return root
		}
	}

	if root.Left != nil && root.Right != nil {
		if (hasChildren(root.Left, p) && hasChildren(root.Right, q)) || (hasChildren(root.Right, p) && hasChildren(root.Left, q)) {
			return root
		}
	}

	if root.Left != nil {
		if hasChildren(root.Left, p) && hasChildren(root.Left, q) {
			return lowestCommonAncestor(root.Left, p, q)
		}
	}

	if root.Right != nil {
		if hasChildren(root.Right, p) && hasChildren(root.Right, q) {
			return lowestCommonAncestor(root.Right, p, q)
		}
	}

	return &TreeNode{}
}

func hasChildren(root, child *TreeNode) bool {
	if root == nil {
		return false
	}
	if root == child {
		return true
	}
	if root.Left != nil && root.Right != nil {
		return hasChildren(root.Left, child) || hasChildren(root.Right, child)
	}
	if root.Left != nil {
		return hasChildren(root.Left, child)
	}
	return hasChildren(root.Right, child)
}
