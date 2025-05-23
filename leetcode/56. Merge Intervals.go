package leetcode

import (
	"log"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}

	start, end := intervals[0][0], intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		newStart, newEnd := intervals[i][0], intervals[i][1]
		if newStart > end {
			res = append(res, []int{start, end})
			start, end = newStart, newEnd
			continue
		}
		if newEnd > end {
			end = newEnd
		}
	}
	res = append(res, []int{start, end})

	return res
}

func merge3(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	n := len(intervals)
	start := []int{}
	end := []int{}

	for _, interval := range intervals {
		start = append(start, interval[0])
		end = append(end, interval[1])
	}

	sort.Ints(start)
	sort.Ints(end)

	begin := start[0]
	res := [][]int{}
	for i := range n - 1 {
		if end[i] < start[i+1] {
			res = append(res, []int{begin, end[i]})
			begin = start[i+1]
		}
	}

	if end[n-2] < start[n-1] {
		res = append(res, []int{start[n-1], end[n-1]})
	} else {
		res = append(res, []int{begin, end[n-1]})
	}

	return res
}

func merge2(intervals [][]int) [][]int {
	res := [][]int{}
	merged := make([]int, 10001)
	duplicate := make(map[int][]byte)

	for _, interval := range intervals {
		if interval[0] == interval[1] {
			if _, ok := duplicate[interval[0]]; !ok {
				duplicate[interval[0]] = []byte{}
			}
			continue
		}

		merged[interval[0]]++
		merged[interval[1]]--
	}

	start := -1
	prefixSum := 0
	for i := 0; i < len(merged); i++ {
		prefixSum += merged[i]

		if start == -1 {
			if prefixSum > 0 {
				start = i
			} else {
				if _, ok := duplicate[i]; ok {
					res = append(res, []int{i, i})
				}
			}

			continue
		}

		if prefixSum == 0 {
			res = append(res, []int{start, i})
			start = -1
		}
	}

	return res
}

func Testmerge() {
	type Case struct {
		intervals [][]int

		expected [][]int
	}

	cases := []Case{
		{
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			intervals: [][]int{{1, 4}, {4, 5}},
			expected:  [][]int{{1, 5}},
		},
		{
			intervals: [][]int{{1, 4}, {0, 0}},
			expected:  [][]int{{0, 0}, {1, 4}},
		},
	}

	for _, testcase := range cases {
		res := merge(testcase.intervals)

		if len(res) != len(testcase.expected) {
			log.Panicf("FAILED. %+v Expected %+v but got %+v", testcase.intervals, testcase.expected, res)
		}
		log.Printf("MATCHED. %+v result: %+v, expected: %+v\n\n", testcase.intervals, res, testcase.expected)
	}

}
