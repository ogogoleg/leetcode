package problems

// 1 2 2

//1 2 

func removeDuplicates(nums []int) int {

	last_unique_element := nums[0] // 1
	next_free_slot_index := 1 

    for i := 1;  i < len(nums); i++ {
		if nums[i] != last_unique_element {
			if next_free_slot_index != i {
				nums[next_free_slot_index] = nums[i]
			} 
			last_unique_element = nums[i]
			next_free_slot_index++
		}
	}
	
	return next_free_slot_index 
}