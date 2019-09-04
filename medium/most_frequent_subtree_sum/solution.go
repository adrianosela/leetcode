package most_frequent_subtree_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	m := make(map[int]int)
	subtreeSum(m, root)
	return maxValKey(m)
}

func maxValKey(m map[int]int) []int {
	keys := []int{}
	max := 0

	for k, v := range m {
		if v == max {
			keys = append(keys, k)
		}
		if v > max {
			keys = []int{k}
			max = v
		}
	}

	return keys
}

func subtreeSum(m map[int]int, r *TreeNode) int {
	var sum int

	// store sum before return
	defer func() {
		if _, ok := m[sum]; ok {
			m[sum]++
		} else {
			m[sum] = 1
		}
	}()

	if r.Left != nil && r.Right != nil {
		sum = r.Val + subtreeSum(m, r.Left) + subtreeSum(m, r.Right)
		return sum
	}
	if r.Left != nil {
		sum = r.Val + subtreeSum(m, r.Left)
		return sum
	}
	if r.Right != nil {
		sum = r.Val + subtreeSum(m, r.Right)
		return sum
	}
	sum = r.Val
	return sum
}
