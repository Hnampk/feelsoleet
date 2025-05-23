package main

import (
	"feelsoleet/sorting"
	"log"
)

func main() {
	algo()
	// 28. Find the Index of the First Occurrence in a String
	// leetcode.TeststrStr()

	// 125. Valid Palindrome
	// leetcode.TestisPalindrome()

	// 392. Is Subsequence
	// leetcode.TestisSubsequence()

	// 167. Two Sum II - Input Array Is Sorted
	// leetcode.TesttwoSum()

	// 11. Container With Most Water
	// leetcode.TestmaxArea()

	// 15. 3Sum
	// leetcode.TestthreeSum()

	// 383. Ransom Note
	// leetcode.TestcanConstruct()

	// 205. Isomorphic Strings
	// leetcode.TestisIsomorphic()

	// 290. Word Pattern
	// leetcode.TestwordPattern()

	// 242. Valid Anagram
	// leetcode.TestisAnagram()

	// 49. Group Anagrams
	// leetcode.TestgroupAnagrams()

	// 202. Happy Number
	// leetcode.TestisHappy()

	// 219. Contains Duplicate II
	// leetcode.TestcontainsNearbyDuplicate()

	// 20. Valid Parentheses
	// leetcode.TestisValid()

	// 128. Longest Consecutive Sequence
	// leetcode.TestlongestConsecutive()

	// 3531. Count Covered Buildings
	// leetcode.TestcountCoveredBuildings()

	// 71. Simplify Path
	// leetcode.TestsimplifyPath()

	// 155. Min Stack

	// 150. Evaluate Reverse Polish Notation
	// leetcode.TestevalRPN()

	// 141. Linked List Cycle
	// leetcode.TesthasCycle()

	// 2. Add Two Numbers
	// leetcode.TestaddTwoNumbers()

	// 373. Find K Pairs with Smallest Sums
	// leetcode.TestkSmallestPairs()

	// 35. Search Insert Position
	// leetcode.TestsearchInsert()

	// 74. Search a 2D Matrix
	// leetcode.TestsearchMatrix()

	// 200. Number of Islands
	// leetcode.TestnumIslands()

	// 130. Surrounded Regions
	// leetcode.Testsolve()

	// 133. Clone Graph
	// leetcode.TestcloneGraph()

	// 399. Evaluate Division
	// leetcode.TestcalcEquation()

	// 162. Find Peak Element
	// leetcode.TestfindPeakElement()

	// 33. Search in Rotated Sorted Array
	// leetcode.Testsearch()

	// 228. Summary Ranges
	// leetcode.TestsummaryRanges()
}

func todo() {
	// 224. Basic Calculator
	// leetcode.Testcalculate() // todo
}

func algo() {
	log.Printf("===== Start Algo =====")
	nums := []int{4, 1, 3, 1, 5, 2}
	log.Printf("before SelectionSort: %+v", nums)
	sorting.SelectionSort(nums)
	log.Printf("after SelectionSort: %+v", nums)

	log.Printf("======================")
	nums = []int{4, 1, 3, 1, 5, 2}
	log.Printf("before BubbleSort: %+v", nums)
	sorting.BubbleSort(nums)
	log.Printf("after BubbleSort: %+v", nums)

	log.Printf("======================")
	nums = []int{4, 1, 3, 1, 5, 2}
	log.Printf("before InsertionSort: %+v", nums)
	sorting.InsertionSort(nums)
	log.Printf("after InsertionSort: %+v", nums)
}
