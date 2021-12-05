package functions

import (
	"fmt"
)

// "Data": Struct for Data from txt file.
type Data struct {
	Ants  int
	Start []string
	End   string
	Links []string

	CheckForString string
}

// "Graph":
type Graph struct {
	vertices []*Vertex
}

// "Vertex":
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// "AddVertex":
func (g *Graph) AddVertex(k int) {
	if Contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

// "Print":
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("%v ", v.key)
		}
	}
	fmt.Println()
}

// "Contains":
func Contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

// "GetVertex":
func (g *Graph) GetVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

// "AddEdge":
func (g *Graph) AddEdge(from, to int) {
	//get vertex
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)
	//check errors
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if Contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("Existing edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else {
		//add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}

}
