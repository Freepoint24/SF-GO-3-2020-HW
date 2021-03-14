package main

// IsBracketsBalanced проверяет сбалансированы ли скобки в строке
// Если строки не сбалансированы, возвращает false и позицию,
// в которой строка несбалансирована
func isBracketsBalanced(str string) (bool, int) {
	var (
		pos   int
		val   rune
		stack = make([]rune, 0)
	)

	for pos, val = range str {
		switch val {
		case '(', '[', '{':
			stack = append(stack, val)
		case ')', ']', '}':
			if len(stack) == 0 {
				return false, pos
			}
			pop := stack[len(stack)-1]
			if (pop == '(' && val != ')') || (pop == '[' && val != ']') || (pop == '{' && val != '}') {
				return false, pos
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0, pos
}
