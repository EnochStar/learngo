package recursion

func generateParenthesis(n int) (res []string) {
	var generate func(left, right int, s string, n int)
	generate = func(left, right int, s string, n int) {
		if left == n && right == n {
			res = append(res, s)
		}
		if left < n {
			generate(left+1, right, s+"(", n)
		}
		if right < left {
			generate(left, right+1, s+")", n)
		}
	}
	generate(0, 0, "", n)
	return res
}
