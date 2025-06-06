package leetcode

import (
	"fmt"
	"log"
)

func readBinaryWatch(turnedOn int) []string {
	res := []string{}
	binaryWatchBacktrack(turnedOn, 0, 0, 0, 0, &res)
	return res
}

const numberOfLeds = 10 // 4 for hours + 6 for minutes
func binaryWatchBacktrack(turnedOn, hours, minutes, curLedIndex, turnedOnByHM int, res *[]string) {
	// 1. record if it's the solution
	if turnedOn == turnedOnByHM {
		minuteStr := fmt.Sprintf("%d", minutes)
		if minutes < 10 {
			minuteStr = "0" + minuteStr
		}
		*res = append(*res, fmt.Sprintf("%d:%s", hours, minuteStr))
		return
	}

	for i := curLedIndex; i < numberOfLeds; i++ {
		// prune?
		if i < 6 { // minutes
			// 2. make choice
			minutes += (1 << i)
			if minutes < 60 { // invalid
				// 3. backtrack
				binaryWatchBacktrack(turnedOn, hours, minutes, i+1, turnedOnByHM+1, res)
			}
			// 4. undo choice
			minutes -= (1 << i)
		} else { // hours
			// 2. make choice
			hours += (1 << (i - 6))
			if hours < 12 { // invalid
				// 3. backtrack
				binaryWatchBacktrack(turnedOn, hours, minutes, i+1, turnedOnByHM+1, res)
			}
			// 4. undo choice
			hours -= (1 << (i - 6))
		}
	}
}

func TestreadBinaryWatch() {
	type Case struct {
		turnedOn int

		expected []string
	}

	cases := []Case{
		{
			turnedOn: 1,
			expected: []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"},
		},
		{
			turnedOn: 9,
			expected: []string{},
		},
	}

	for _, testcase := range cases {
		res := readBinaryWatch(testcase.turnedOn)

		// if res != testcase.expected {
		// 	log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.numbers, testcase.target, testcase.expected, res)
		// }
		if len(res) != len(testcase.expected) {
			log.Panicf("FAILED. %+v, Expected %+v but got %+v", testcase.turnedOn, testcase.expected, res)
		}
		for i := 0; i < len(testcase.expected); i++ {
			if testcase.expected[i] != res[i] {
				log.Panicf("FAILED. %+v, Expected %+v but got %+v", testcase.turnedOn, testcase.expected, res)
			}
		}

		log.Printf("MATCHED. %+v, result: %+v, expected: %+v\n\n", testcase.turnedOn, res, testcase.expected)
	}
}
