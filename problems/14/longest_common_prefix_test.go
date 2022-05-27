// Example 1:

// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:

// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.
package problems

import (
	"testing"

	"gotest.tools/assert"
)

 func TestLongestCommonPrefix(t *testing.T) {
	 result := longestCommonPrefix([]string{"flower","flow","flight"})
	 assert.Equal(t, result, "fl")

	 result = longestCommonPrefix([]string{"dog","racecar","car"})
	 assert.Equal(t, result, "")

	 result = longestCommonPrefix([]string{"flower","","flight"})
	 assert.Equal(t, result, "")

	 result = longestCommonPrefix([]string{"","flow","flight"})
	 assert.Equal(t, result, "")

	 result = longestCommonPrefix([]string{""})
	 assert.Equal(t, result, "")

	 result = longestCommonPrefix([]string{"abc"})
	 assert.Equal(t, result, "abc")
 }