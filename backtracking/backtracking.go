package backtracking

func backtrack(state *State, choices []Choice, res *[]State) {
	// Check if it's a solution
	if isSolution(state) {
		// Record the solution
		recordSolution(state, res)
		// Stop searching
		return
	}
	// Iterate through all choices
	for _, choice := range choices {
		// Prune: check if the choice is valid
		if isValid(state, choice) {
			// Trial: make a choice, update the state
			makeChoice(state, choice)
			backtrack(state, choices, res)
			// Retreat: undo the choice, revert to the previous state
			undoChoice(state, choice)
		}
	}
}

type State struct {
	Path []Choice
}

type Choice struct {
	// tree node
	Val int
}

func recordSolution(state *State, res *[]State) {
	*res = append(*res, *state)
}

func isSolution(state *State) bool {
	if len(state.Path) == 0 {
		return false
	}
	return state.Path[len(state.Path)-1].Val == 7
}

func isValid(state *State, choice Choice) bool {
	return &choice != nil && choice.Val != 3

}
func makeChoice(state *State, choice Choice) {
	state.Path = append(state.Path, choice)
}
func undoChoice(state *State, choice Choice) {
	state.Path = state.Path[1:]
}
