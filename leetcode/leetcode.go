package main

import "math"

func main() {

}
func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return res
}
func bag(num, weight int, w, v []int) int {
	dp := make([][]int, num+1, weight+1)
	for i := 0; i <= num; i++ {
		dp[i][0] = 0
		dp[0][i] = 0
	}
	for i := 1; i <= num; i++ {
		for j := 1; j <= weight; j++ {
			if w[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i-1][j-w[i-1]]+v[i-1])))
			}
		}
	}
	return dp[num][weight]
}
func lengthOfLongestSubstring(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
