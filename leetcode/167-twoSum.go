package leetcode

import "log"

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		if numbers[left]+numbers[right] > target {
			right--
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}

	return []int{}
}
func twoSum2(numbers []int, target int) []int {
	// declare an map to manage index of value is numbers
	// loop through numbers, calculate the SUB = target - number[i], and find the index of SUB in map.
	//	return if found
	//	of add info of number[i] to map

	mapData := make(map[int]int)

	for i, number := range numbers {
		if found, ok := mapData[target-number]; ok {
			if i < found {
				return []int{i + 1, found + 1}
			}
			return []int{found + 1, i + 1}
		} else {
			mapData[number] = i
		}
	}

	return []int{}
}

func twoSum1(numbers []int, target int) []int {
	// BeginFindingIndex = len(numbers) - 1
	// iterate through numbers[i], calculate SUB = target - numbers[i].
	// iterate through numbers[BeginFindingIndex] until BeginFindingIndex = i+1. Return if found numbers[i] == SUB.
	//	Or break when numbers[i] < SUB and marks BeginFindingIndex = i

	BeginFindingIndex := len(numbers) - 1
	for i := 0; i < len(numbers)-1; i++ {

		for j := BeginFindingIndex; j > i; j-- {
			if numbers[j] == target-numbers[i] {
				return []int{i + 1, j + 1}
			}

			if numbers[j] < target-numbers[i] {
				BeginFindingIndex = j
				break
			}
		}
	}

	return []int{}
}

func TesttwoSum() {
	type Case struct {
		numbers []int
		target  int

		expected []int
	}

	cases := []Case{
		{
			numbers:  []int{2, 7, 11, 15},
			target:   9,
			expected: []int{1, 2},
		},
		{
			numbers:  []int{2, 3, 4},
			target:   6,
			expected: []int{1, 3},
		},
		{
			numbers:  []int{-1, 0},
			target:   -1,
			expected: []int{1, 2},
		},
		{
			numbers:  []int{5, 25, 75},
			target:   100,
			expected: []int{2, 3},
		},
	}

	for _, testcase := range cases {
		res := twoSum(testcase.numbers, testcase.target)

		if len(res) != 2 {
			log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.numbers, testcase.target, testcase.expected, res)
		}
		for i := 0; i < len(testcase.expected); i++ {
			if testcase.expected[i] != res[i] {
				log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.numbers, testcase.target, testcase.expected, res)
			}
		}

		log.Printf("MATCHED. %+v, %+v, result: %+v, expected: %+v\n\n", testcase.numbers, testcase.target, res, testcase.expected)
	}
}
