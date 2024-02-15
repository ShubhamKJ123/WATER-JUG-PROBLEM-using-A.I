package main

import (
	"fmt"
)

type pair struct {
	first  int
	second int
}

func BFS(a, b, target int) {
	m := make(map[pair]bool)
	isSolvable := false
	var path []pair
	q := []pair{{0, 0}}

	fmt.Println("Steps:")
	for len(q) > 0 {
		u := q[0]
		q = q[1:]

		if _, ok := m[u]; ok {
			continue
		}

		if u.first == target || u.second == target {
			isSolvable = true
			path = append(path, u)

			for i := 0; i < len(path); i++ {
				fmt.Println(path[i])
			}
			break
		}

		m[u] = true

		q = append(q, pair{a, u.second}) // Fill jug 1
		fmt.Println("Fill jug 1:", a, u.second)

		q = append(q, pair{u.first, b}) // Fill jug 2
		fmt.Println("Fill jug 2:", u.first, b)

		q = append(q, pair{0, u.second}) // Empty jug 1
		fmt.Println("Empty jug 1:", 0, u.second)

		q = append(q, pair{u.first, 0}) // Empty jug 2
		fmt.Println("Empty jug 2:", u.first, 0)

		if u.second != 0 && u.first != a { // Pour from jug 2 to jug 1
			q = append(q, pair{min(u.first+u.second, a), max(0, u.second-(a-u.first))})
			fmt.Println("Pour water from jug 2 into jug 1:", min(u.first+u.second, a), max(0, u.second-(a-u.first)))
		}

		if u.first != 0 && u.second != b { // Pour from jug 1 to jug 2
			q = append(q, pair{max(0, u.first-(b-u.second)), min(u.second+u.first, b)})
			fmt.Println("Pour water from jug 1 into jug 2:", max(0, u.first-(b-u.second)), min(u.second+u.first, b))
		}
	}

	if !isSolvable {
		fmt.Println("No solution")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	jug1, jug2, target := 4, 3, 2
	fmt.Println("Path from initial state to solution state:")
	BFS(jug1, jug2, target)
}
