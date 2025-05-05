package leetcode

import "log"

func isIsomorphic(s string, t string) bool {
	// loop through each character of both s and t, check if:
	//	character at index i of s, after converted through the hashmap, must match to the character at index i of t
	//	if the character doesn't exists yet, set a new pair into the hashmap
	//	for any mismatch, return false

	s2t := make(map[byte]byte)
	tExisted := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if tfroms, ok := s2t[s[i]]; !ok {
			if _, ok := tExisted[t[i]]; ok {
				// already map to another char
				return false
			}

			s2t[s[i]] = t[i]
			tExisted[t[i]] = byte(0)

		} else {
			if tfroms != t[i] {
				return false
			}
		}
	}

	return true
}

func TestisIsomorphic() {

	type Case struct {
		s string
		t string

		expected bool
	}

	cases := []Case{
		{
			s:        "egg",
			t:        "add",
			expected: true,
		},
		{
			s:        "foo",
			t:        "bar",
			expected: false,
		},
		{
			s:        "paper",
			t:        "title",
			expected: true,
		},
	}

	for _, testcase := range cases {
		res := isIsomorphic(testcase.s, testcase.t)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.s, testcase.t, testcase.expected, res)
		}

		log.Printf("MATCHED. %+v, %+v, result: %+v, expected: %+v\n\n", testcase.s, testcase.t, res, testcase.expected)
	}
}
