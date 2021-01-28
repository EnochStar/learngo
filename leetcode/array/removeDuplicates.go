package array

func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	fis := 0
	las := 1
	for las != len(nums) {
		if nums[fis] != nums[las] {
			fis++
			nums[fis] = nums[las]
		}
		las++
	}
	return fis + 1
}
