package problems

import "math"

var closing map[rune]rune = map[rune]rune{'{':'}', '[':']', '(':')'}
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.

func isValid(s string) bool {

	length := len(s)

    if math.Mod(float64(length),2) != 0 {
		return false
	}

	var processed []rune

	for _, current := range s {
		if current == '(' || current == '{' || current == '[' {
			processed = append(processed, current)
		} else {
			
			if len(processed) == 0 {
				return false
			}

			prev := processed[len(processed)-1]
			if (closing[prev] != current) {
				return false
			}
			processed = processed[:len(processed)-1]
		}
	}

	return len(processed) == 0
}

