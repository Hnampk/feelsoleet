package main

import (
	"feelsoleet/contest"
	"feelsoleet/leetcode"
	"feelsoleet/sorting"
	"log"
)

func main() {
	// algo()
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

	// 108. Convert Sorted Array to Binary Search Tree
	// leetcode.TestsortedArrayToBST()

	// 226. Invert Binary Tree
	// leetcode.TestinvertTree()

	// 190. Reverse Bits
	// leetcode.TestreverseBits()

	// 191. Number of 1 Bits
	// leetcode.TesthammingWeight()

	// 1763. Longest Nice Substring
	// leetcode.TestlongestNiceSubstring()

	// 427. Construct Quad Tree
	// leetcode.Testconstruct()

	// 101. Symmetric Tree
	// leetcode.TestisSymmetric()

	// 1382. Balance a Binary Search Tree
	// leetcode.TestbalanceBST()

	// 105. Construct Binary Tree from Preorder and Inorder Traversal
	// leetcode.TestbuildTree()

	// 558. Logical OR of Two Binary Grids Represented as Quad-Trees
	// leetcode.Testintersect()

	// 106. Construct Binary Tree from Inorder and Postorder Traversal
	// leetcode.TestbuildTree106()

	// 395. Longest Substring with At Least K Repeating Characters
	// leetcode.TestlongestSubstring()

	// 654. Maximum Binary Tree
	// leetcode.TestconstructMaximumBinaryTree()

	// 117. Populating Next Right Pointers in Each Node II
	// leetcode.Testconnect117()

	// 257. Binary Tree Paths
	// leetcode.TestbinaryTreePaths()

	// 401. Binary Watch
	// leetcode.TestreadBinaryWatch()

	// 46. Permutations
	// leetcode.Testpermute()

	// 1863. Sum of All Subset XOR Totals
	// leetcode.TestsubsetXORSum()

	// 988. Smallest String Starting From Leaf
	// leetcode.TestsmallestFromLeaf()

	// 77. Combinations
	// leetcode.Testcombine()

	// 17. Letter Combinations of a Phone Number
	// leetcode.TestletterCombinations()

	// 51. N-Queens
	// leetcode.TestsolveNQueens()

	// 39. Combination Sum
	// leetcode.TestcombinationSum()

	// 22. Generate Parentheses
	// leetcode.TestgenerateParenthesis()

	// 79. Word Search
	// leetcode.Testexist()

	// 64. Minimum Path Sum
	leetcode.TestminPathSum()
}

func todo() {
	// 224. Basic Calculator
	// leetcode.Testcalculate() // todo

	// 2359. Find Closest Node to Given Two Nodes
	// leetcode.TestclosestMeetingNode()
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

	// 3355. Zero Array Transformation I
	// leetcode.TestisZeroArray()

	// 56. Merge Intervals
	// leetcode.Testmerge()

	// 209. Minimum Size Subarray Sum
	leetcode.TestminSubArrayLen()

	// RunContest()
}

func RunContest() {
	// Q1. Smallest Index With Digit Sum Equal to Index
	contest.TestsmallestIndex()
}

// divide and conquer
// - todo Tower of Hanoi problem
// - todo Large integer multiplication
// - todo binary search using recursion instead of iteration
// - todo build tree https://www.hello-algo.com/en/chapter_divide_and_conquer/build_binary_tree_problem/#4-code-implementation
