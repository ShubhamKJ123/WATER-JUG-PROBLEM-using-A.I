package main

import "fmt"

type pair struct {
	first, second int
}

func BFS(a, b, target int) {
	m := make(map[pair]bool)
	isSolvable := false
	var path []pair
	q := []pair{{0, 0}}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]

		if _, ok := m[u]; ok {
			continue
		}
		if u.first > a || u.second > b || u.first < 0 || u.second < 0 {
			continue
		}

		path = append(path, u)

		m[u] = true

		if u.first == target || u.second == target {
			isSolvable = true

			if u.first == target {
				if u.second != 0 {
					path = append(path, pair{u.first, 0})
				}
			} else {
				if u.first != 0 {
					path = append(path, pair{u.second, 0})
				}
			}

			for i := 0; i < len(path); i++ {
				fmt.Printf("(%d, %d)\n", path[i].first, path[i].second)
			}
			break
		}

		q = append(q, pair{u.first, b})
		q = append(q, pair{a, u.second})

		for ap := 0; ap <= max(a, b); ap++ {
			c := u.first + ap
			d := u.second - ap

			if c == a || (d == 0 && d >= 0) {
				q = append(q, pair{c, d})
			}

			c = u.first - ap
			d = u.second + ap

			if (c == 0 && c >= 0) || d == b {
				q = append(q, pair{c, d})
			}
		}

		q = append(q, pair{a, 0})
		q = append(q, pair{0, b})
	}

	if !isSolvable {
		fmt.Println("No solution")
	}
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
