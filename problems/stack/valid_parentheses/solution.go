package valid_parentheses

import ds "learning/internal/ds/stack"

func ValidParentheses(s string) bool {
	parMap := map[rune]rune{'(': ')', '[': ']', '{': '}'}

	stack := ds.Stack{}

	for _, par := range s {
		if aPar, ok := parMap[par]; ok {
			stack.Push(int(aPar))
		} else {
			if stack.Length() == 0 {
				return false
			}
			val, _ := stack.Pop()
			if val != int(par) {
				return false
			}
		}
	}
	if stack.Length() > 0 {
		return false
	}
	return true
}
