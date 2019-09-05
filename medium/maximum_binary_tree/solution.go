package maximum_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	// is leaf
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	// not a leaf, find max index
	max := 0
	var maxIndex int
	for i, elem := range nums {
		if elem > max {
			maxIndex = i
			max = elem
		}
	}

	// exception to the rule - root is at first index
	if maxIndex == 0 {
		return &TreeNode{
			Val:   max,
			Right: constructMaximumBinaryTree(nums[maxIndex+1:]),
		}
	}

	// exception to the rule - root is at last index
	if maxIndex == len(nums)-1 {
		return &TreeNode{
			Val:  max,
			Left: constructMaximumBinaryTree(nums[:maxIndex]),
		}
	}

	return &TreeNode{
		Val:   max,
		Left:  constructMaximumBinaryTree(nums[:maxIndex]),
		Right: constructMaximumBinaryTree(nums[maxIndex+1:]),
	}
}
