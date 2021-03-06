// https://leetcode-cn.com/problems/wiggle-subsequence/
// 376. 摆动序列

package wigglemaxlength

// 贪心
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	ans := 1
	prevDiff := nums[1] - nums[0]
	if prevDiff != 0 {
		ans = 2
	}
	for i := 2; i < n; i++ {
		diff := nums[i] - nums[i-1]
		if diff > 0 && prevDiff <= 0 || diff < 0 && prevDiff >= 0 {
			ans++
			prevDiff = diff
		}
	}
	return ans
}

// dp
// func wiggleMaxLength(nums []int) int {
// 	n := len(nums)
// 	if n < 2 {
// 		return n
// 	}
// 	up, down := 1, 1
// 	for i := 1; i < n; i++ {
// 		if nums[i] > nums[i-1] {
// 			up = down + 1
// 		} else if nums[i] < nums[i-1] {
// 			down = up + 1
// 		}
// 	}
// 	return max(up, down)
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// 作者：LeetCode-Solution
// 链接：https://leetcode-cn.com/problems/wiggle-subsequence/solution/bai-dong-xu-lie-by-leetcode-solution-yh2m/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
