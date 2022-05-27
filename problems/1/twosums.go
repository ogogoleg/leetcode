package problems

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

func TwoSums(nums []int, target int) []int {

	elements := make(map[int]int)

	for i, v := range nums {
		elem, ok := elements[target-v]
		if ok {
			return []int{elem, i}
		}
		elements[v] = i
	}
	return []int{}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
