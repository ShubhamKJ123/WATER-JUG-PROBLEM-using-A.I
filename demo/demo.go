package main

import "fmt"

type Graph struct {
	verices []*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

func (g *Graph) AddVertex(k int) {
	g.verices = append(g.verices, &Vertex{key: k})
}

// func (g *Graph) getVertex(k int) *Vertex {
// 	for i, v := range g.verices {
// 		if v.key == k {
// 			return g.vertices[i]
// 		}
// 	}
// 	return nil
// }

func main() {
	test := &Graph{
		verices: []*Vertex{},
	}

	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}

	fmt.Println(test)
}
