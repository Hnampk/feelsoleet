package backtracking

import "sort"

var (
	duplicateSubsetPruning = false // [4,5] & [5,4] are equal
)

func TestbacktrackSubsetSum() {
	state := &StateSubsetSum{
		data: []ChoiceSubsetSum{},
	}

	target := 9
	choices := []ChoiceSubsetSum{
		{3}, {4}, {5},
	}

	res := []StateSubsetSum{}

	if duplicateSubsetPruning {
		// sort first
		sort.Slice(choices, func(i, j int) bool {
			return choices[i].Val > choices[j].Val
		})
	}

	backtrackSubsetSum(state, target, choices, &res)
}

func backtrackSubsetSum(state *StateSubsetSum, target int, choices []ChoiceSubsetSum, res *[]StateSubsetSum) {
	// Check if it's a solution
	if isSolutionSubsetSum(state, target) {
		// Record the solution
		recordSolutionSubsetSum(state, res)
		// Stop searching
		return
	}
	// Iterate through all choices
	for _, choice := range choices {
		// Prune: check if the choice is valid
		if isValidSubsetSum(state, choice, target) {
			// Trial: make a choice, update the state
			makeChoiceSubsetSum(state, choice)
			backtrackSubsetSum(state, target, choices[1:], res)
			// Retreat: undo the choice, revert to the previous state
			undoChoiceSubsetSum(state, choice)
		}
	}
}

type StateSubsetSum struct {
	data []ChoiceSubsetSum
}

type ChoiceSubsetSum struct {
	Val int
}

func (s *StateSubsetSum) total() int {
	total := 0
	for _, item := range s.data {
		total += item.Val
	}

	return total
}

func isSolutionSubsetSum(state *StateSubsetSum, target int) bool {
	return state.total() == target
}

func recordSolutionSubsetSum(state *StateSubsetSum, res *[]StateSubsetSum) {
	*res = append(*res, *state)
}

func isValidSubsetSum(state *StateSubsetSum, choice ChoiceSubsetSum, target int) bool {
	return state.total() <= target
}

func makeChoiceSubsetSum(state *StateSubsetSum, choice ChoiceSubsetSum) {
	state.data = append(state.data, choice)
}

func undoChoiceSubsetSum(state *StateSubsetSum, choice ChoiceSubsetSum) {
	state.data = state.data[:len(state.data)-1] // pop
}
