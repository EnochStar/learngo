package dp

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	x := 2
	y := 3
	var sum int
	for i := 4; i <= n; i++ {
		sum = x + y
		x = y
		y = sum
	}
	return sum
}
