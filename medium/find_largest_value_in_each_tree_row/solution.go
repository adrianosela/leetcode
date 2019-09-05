package find_largest_value_in_each_tree_row

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	// safety check
	if root == nil {
		return []int{}
	}

	// is leaf
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	// has only one
	if root.Left == nil {
		return append([]int{root.Val}, largestValues(root.Right)...)
	}
	if root.Right == nil {
		return append([]int{root.Val}, largestValues(root.Left)...)
	}

	// has both
	ml := largestValues(root.Left)
	mr := largestValues(root.Right)
	ll := len(ml)
	lr := len(mr)

	if ll > lr {
		mm := make([]int, ll)
		i := 0
		for i < lr {
			mm[i] = int(math.Max(float64(ml[i]), float64(mr[i])))
			i++
		}
		for i < ll {
			mm[i] = ml[i]
			i++
		}
		return append([]int{root.Val}, mm...)
	} else if lr > ll {
		mm := make([]int, lr)
		i := 0
		for i < ll {
			mm[i] = int(math.Max(float64(ml[i]), float64(mr[i])))
			i++
		}
		for i < lr {
			mm[i] = mr[i]
			i++
		}
		return append([]int{root.Val}, mm...)
	} else {
		mm := make([]int, len(ml)) // same size
		for i := 0; i < len(mr); i++ {
			mm[i] = int(math.Max(float64(ml[i]), float64(mr[i])))
		}
		return append([]int{root.Val}, mm...)
	}
}
