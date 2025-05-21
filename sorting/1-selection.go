package sorting

// SelectionSort has time complexity O(n^2)
func SelectionSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		k := i

		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[k] {
				k = j
			}
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
}
