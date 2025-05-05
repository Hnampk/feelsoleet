package leetcode

import (
	"log"
	"strings"
)

func simplifyPath(path string) string {
	dirs := []string{}

	startAt := 0
	for i := 0; i < len(path); i++ {
		if path[i] == '/' {
			if startAt == i {
				startAt = i + 1
				continue
			}

			dirs = appendDir(dirs, path, startAt, i)

			startAt = i + 1
		} else if i == len(path)-1 {
			dirs = appendDir(dirs, path, startAt, i+1)
		}
	}

	return "/" + strings.Join(dirs, "/")
}

func appendDir(dirs []string, path string, start, to int) []string {
	if path[start:to] == ".." {
		if len(dirs) > 0 {
			dirs = dirs[:len(dirs)-1]
		}
	} else if path[start:to] != "." {
		dirs = append(dirs, path[start:to])
	}
	return dirs
}

func simplifyPath2(path string) string {
	dir := []string{}
	value := ""

	for i := 1; i < len(path); i++ {
		if path[i] != '/' {
			value += string(path[i])
		}

		if path[i] == '/' || i == len(path)-1 {
			if value == ".." {
				if len(dir) > 0 {
					dir = dir[:len(dir)-1] // pop
				}
			} else if value != "" && value != "." {
				dir = append(dir, value)
			}

			value = ""
			continue
		}
	}

	return "/" + strings.Join(dir, "/")
}

func TestsimplifyPath() {
	type Case struct {
		path string

		expected string
	}

	cases := []Case{
		{
			path:     "/home/",
			expected: "/home",
		},
		{
			path:     "/home//foo/",
			expected: "/home/foo",
		},
		{
			path:     "/home/user/Documents/../Pictures",
			expected: "/home/user/Pictures",
		},
		{
			path:     "/../",
			expected: "/",
		},
		{
			path:     "/./",
			expected: "/",
		},
		{
			path:     "/.../a/../b/c/../d/./",
			expected: "/.../b/d",
		},
	}

	for _, testcase := range cases {
		res := simplifyPath(testcase.path)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v Expected %+v but got %+v", testcase.path, testcase.expected, res)
		}

		log.Printf("MATCHED. %+v result: %+v, expected: %+v\n\n", testcase.path, res, testcase.expected)
	}

}
