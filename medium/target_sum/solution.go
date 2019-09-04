package target_sum

func findTargetSumWays(nums []int, S int) int {
	if len(nums) == 0 {
		if S == 0 {
			return 1
		}
		return 0
	}
	return findTargetSumWays(nums[1:], S-nums[0]) + findTargetSumWays(nums[1:], S+nums[0])
}
