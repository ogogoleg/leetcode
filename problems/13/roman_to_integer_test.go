package problems

import (
	"testing"

	"gotest.tools/assert"
)

func TestRomanToInt(t *testing.T) {

	assert.Equal(t, RomanToInt("III"), 3)
	// Input: s = "III"
	// Output: 3
	// Explanation: III = 3.

	assert.Equal(t, RomanToInt("LVIII"), 58)
	// Input: s = "LVIII"
	// Output: 58
	// Explanation: L = 50, V= 5, III = 3.
	assert.Equal(t, RomanToInt("MCMXCIV"), 1994)
	// Input: s = "MCMXCIV"
	// Output: 1994
	// Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

}
