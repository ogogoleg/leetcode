package problems

import (
	"testing"
)

func TestTwoSums(t *testing.T) {

	// 	Input: nums = [2,7,11,15], target = 9
	// Output: [0,1]
	// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
	// Example 2:

	// Input: nums = [3,2,4], target = 6
	// Output: [1,2]
	// Example 3:

	// Input: nums = [3,3], target = 6
	// Output: [0,1]

	//t.("Failing test")

	res := TwoSums([]int{2, 7, 11, 15}, 9)
	if !Equal(res, []int{0, 1}) {
		t.Error()
	}

	res = TwoSums([]int{3, 2, 4}, 6)
	if !Equal(res, []int{1, 2}) {
		t.Error()
	}

	res = TwoSums([]int{3, 3}, 6)
	if !Equal(res, []int{0, 1}) {
		t.Error()
	}

}
