package main

import "fmt"

type state struct {
	m1, c1, b, m2, c2 int
}

func isValid(s state) bool {
	return s.m1 >= 0 && s.m1 <= 3 && s.c1 >= 0 && s.c1 <= 3 &&
		s.m2 >= 0 && s.m2 <= 3 && s.c2 >= 0 && s.c2 <= 3 &&
		((s.m1 == 0 || s.m1 >= s.c1) && (s.m2 == 0 || s.m2 >= s.c2))
}

func generateNextStates(s state) []state {
	moves := [][2]int{{1, 0}, {2, 0}, {0, 1}, {0, 2}, {1, 1}}
	var nextStates []state
	for _, move := range moves {
		dm, dc := move[0], move[1]
		if s.b == 1 {
			newState := state{s.m1 - dm, s.c1 - dc, 0, s.m2 + dm, s.c2 + dc}
			if isValid(newState) {
				nextStates = append(nextStates, newState)
			}
		} else {
			newState := state{s.m1 + dm, s.c1 + dc, 1, s.m2 - dm, s.c2 - dc}
			if isValid(newState) {
				nextStates = append(nextStates, newState)
			}
		}
	}
	return nextStates
}

func solveBFS(start state) []state {
	queue := []struct {
		state
		path []state
	}{{start, []state{start}}}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		currentState, path := current.state, current.path
		if currentState == (state{0, 0, 0, 3, 3}) {
			return path
		}
		for _, nextState := range generateNextStates(currentState) {
			if !contains(path, nextState) {
				newPath := append(path, nextState)
				queue = append(queue, struct {
					state
					path []state
				}{nextState, newPath})
			}
		}
	}
	return nil
}

func contains(path []state, s state) bool {
	for _, p := range path {
		if p == s {
			return true
		}
	}
	return false
}

func main() {
	startState := state{3, 3, 1, 0, 0}
	solution := solveBFS(startState)
	if solution != nil {
		fmt.Println("solution paths:")
		for _, step := range solution {
			fmt.Printf("%d missionaries, %d cannibals, boat on %s, %d missionaries, %d cannibals\n",
				step.m1, step.c1, map[bool]string{true: "left", false: "right"}[step.b == 1], step.m2, step.c2)
		}
	} else {
		fmt.Println("No solution")
	}
}
