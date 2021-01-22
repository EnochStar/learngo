package recursionClosure

func Factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * Factorial(i-1)
}

func Fibonacci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {

		return 1

	}
	return Fibonacci(i-2) + Fibonacci(i-1)
}

func FibonacciPre(i int) int {
	if i <= 1 {
		return i
	}
	pre := 0
	cur := 1
	sum := 0
	for j := 2; j <= i; j++ {
		sum = pre + cur
		pre = cur
		cur = sum
	}
	return sum
}
