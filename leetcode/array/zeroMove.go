package array

func moveZeroes(nums []int) {
	insertIdx := 0
	len := len(nums)
	for i := 0; i < len; i++ {
		if nums[i] != 0 {
			nums[insertIdx] = nums[i]
			insertIdx++
		}
	}
	for insertIdx < len {
		nums[insertIdx] = 0
		insertIdx++
	}
}
func moveZeroesSec(nums []int) {
	j := 0
	len := len(nums)
	for i := 0; i < len; i++ {
		if nums[i] != 0 {
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp
			j++
		}
	}
}
