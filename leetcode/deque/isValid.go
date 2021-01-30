package deque

func isValid(s string) bool {
	var stack []byte
	var top int = -1
	for _, v := range s {
		if v == '(' {
			stack = append(stack, ')')
			top++
		} else if v == '[' {
			stack = append(stack, ']')
			top++
		} else if v == '{' {
			stack = append(stack, '}')
			top++
		} else {
			if top > -1 && stack[top] == byte(v) {
				stack = stack[:top]
				top--
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
