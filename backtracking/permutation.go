package backtracking

var (
	selected          = make([]int, 3) // 3 is the len of choices
	considerDuplicate = false
)

func backtrackPermutation(state *StatePermutation, choices []ChoicePermutation, res *[]StatePermutation) {
	// Check if it's a solution
	if isSolutionPermutation(state) {
		// Record the solution
		recordSolutionPermutation(state, res)
		// Stop searching
		return
	}
	// Iterate through all choices
	duplicated := make(map[int][]byte) // choice data duplicate
	for i, choice := range choices {
		if considerDuplicate {
			if _, ok := duplicated[choice.data]; ok {
				continue
			}
			duplicated[choice.data] = []byte{}
		}
		// Prune: check if the choice is valid
		if isValidPermutationPermutation(state, choice, i) {
			selected[i] = 1
			// Trial: make a choice, update the state
			makeChoicePermutation(state, choice)
			backtrackPermutation(state, choices, res)
			// Retreat: undo the choice, revert to the previous state
			undoChoicePermutation(state, choice)
			selected[i] = 0
		}
	}
}

type StatePermutation struct {
	data []ChoicePermutation
}
type ChoicePermutation struct {
	data int
}

func isSolutionPermutation(state *StatePermutation) bool {
	return len(state.data) == 3
}
func recordSolutionPermutation(state *StatePermutation, res *[]StatePermutation) {
	*res = append(*res, *state)
}
func isValidPermutationPermutation(state *StatePermutation, choice ChoicePermutation, index int) bool {
	return selected[index] == 0
}
func makeChoicePermutation(state *StatePermutation, choice ChoicePermutation) {
	state.data = append(state.data, choice)
}
func undoChoicePermutation(state *StatePermutation, choice ChoicePermutation) {
	state.data = state.data[:len(state.data)-1] // pop
}
