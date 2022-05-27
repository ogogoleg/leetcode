package problems

import (
	"testing"

	"gotest.tools/assert"
)

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.

func TestIsValid(t *testing.T) {

	// Input: s = "()"
	// Output: true

	// Input: s = "()[]{}"
	// Output: true

	// Input: s = "(]"
	// Output: false

	assert.Assert(t, isValid("()"))
	assert.Assert(t, !isValid("(("))
	assert.Assert(t, !isValid("))"))
	assert.Assert(t, !isValid("){"))
	assert.Assert(t, !isValid(")("))
	assert.Assert(t, !isValid("(]"))
	assert.Assert(t, isValid("()[]{}"))
	assert.Assert(t, isValid("(())"))
	assert.Assert(t, isValid("([])"))
	assert.Assert(t, !isValid("([})"))
}