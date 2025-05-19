package main

import "log"

func insertingSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		min := nums[i]
		k := i

		for j := i + 1; j < len(nums); j++ {
			if nums[j] < min {
				min = nums[j]
				k = j
			}
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
}

func TestinsertingSort() {
	nums := []int{1, 7, 2, 3, 5, 4}
	insertingSort(nums)
	log.Println(nums)
}
