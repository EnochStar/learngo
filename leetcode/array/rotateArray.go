package array

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(0, len(nums)-1, nums)
	reverse(0, k-1, nums)
	reverse(k, len(nums)-1, nums)
}
func reverse(start int, end int, nums []int) {
	for start < end {
		tmp := nums[start]
		nums[start] = nums[end]
		nums[end] = tmp
		start++
		end--
	}
}
