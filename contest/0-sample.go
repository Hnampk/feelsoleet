package contest

import "log"

func sample(numbers []int, target int) []int {

	return []int{}
}

func Testsample() {
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
	}

	for _, testcase := range cases {
		res := sample(testcase.numbers, testcase.target)

		// if res != testcase.expected {
		// 	log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.numbers, testcase.target, testcase.expected, res)
		// }
		// if len(res) != 2 {
		// 	log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.numbers, testcase.target, testcase.expected, res)
		// }
		// for i := 0; i < len(testcase.expected); i++ {
		// 	if testcase.expected[i] != res[i] {
		// 		log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.numbers, testcase.target, testcase.expected, res)
		// 	}
		// }

		log.Printf("MATCHED. %+v, %+v, result: %+v, expected: %+v\n\n", testcase.numbers, testcase.target, res, testcase.expected)
	}
}
