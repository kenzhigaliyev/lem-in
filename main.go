package main

import (
	"fmt"
	"student/functions"
)

func main() {

	functions.ReadFile()
	functions.DataForVetices()
	functions.DataForLinks()
	fmt.Println(functions.NewData)
	// fmt.Print(functions.test)
	// test := &functions.Graph{}
	// for i := 0; i < 5; i++ {
	// 	test.AddVertex(i)
	// }

	// test.AddVertex(0)
	// test.AddVertex(0)
	// test.AddEdge(1, 2)
	// test.AddEdge(1, 2)
	// test.AddEdge(6, 2)
	// test.AddEdge(3, 2)
	// test.AddEdge(3, 4)

	// test.Print()

}
