package sorting

// InsertionSort has time complexity O(n^2)
func InsertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		position := i
		for position > 0 {
			if num >= nums[position-1] {
				break
			}

			nums[position] = nums[position-1]
			position--
		}

		nums[position] = num
	}
}
