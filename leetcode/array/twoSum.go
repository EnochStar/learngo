package array

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[x]; ok {
			return []int{i, p}
		}
		hashTable[target-x] = i
	}
	return nil
}
