package functions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// "NewData": Variable for Data{}.
var NewData = &Data{}

var test = &Graph{}

// "FillingData": Reading data from txt file in order to fill the struct.
func FillingData(str string) {
	if str == "##start" {
		NewData.CheckForString = "DataForStart"
		return
	} else if str == "##end" {
		NewData.CheckForString = "DataForEnd"
		return
	}

	if len(NewData.CheckForString) == 0 {
		NewData.Ants, _ = strconv.Atoi(str)
	}

	if NewData.CheckForString == "DataForStart" {
		NewData.Start = append(NewData.Start, str)
	}

	if NewData.CheckForString == "DataForEnd" && !strings.Contains(str, "-") {
		NewData.End = str
	}
	if NewData.CheckForString == "DataForEnd" && strings.Contains(str, "-") {
		NewData.Links = append(NewData.Links, str)
	}
}

// "ReadFile": Function for reading from txt file line by line.
func ReadFile() {
	file, err := os.Open("example1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	for scanner.Scan() {
		FillingData(scanner.Text())
		// fmt.Println()
	}
	NewData.CheckForString = ""
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func DataForVetices() {
	// for _, val := range NewData.Start {
	// 	vertices := strings.Split(val, " ")
	// 	vertice, _ := strconv.Atoi(vertices[0])
	// 	x_coordinate,_ := strconv.Atoi(vertices[1])
	// 	y_coordinate,_ := strconv.Atoi(vertices[2])
	// 	test.AddVertex(vertice)
	// 	test.
	// }

	for i := 0; i < len(NewData.Start); i++ {
		j, _ := strconv.Atoi(strings.Split(NewData.Start[i], " ")[0])
		test.AddVertex(j)
	}
	j, _ := strconv.Atoi(strings.Split(NewData.End, " ")[0])
	test.AddVertex(j)
}

func DataForLinks() {
	for _, val := range NewData.Links {
		links := strings.Split(val, "-")
		from, _ := strconv.Atoi(links[0])
		to, _ := strconv.Atoi(links[1])
		test.AddEdge(from, to)
		test.AddEdge(to, from)
	}

}

// func DFS(g *Graph, startVertex *Vertex, visitCallBack func(int), visited map[int]bool) map[int]bool {

// 	fmt.Println(visited, startVertex.key)

// 	if startVertex == nil {
// 		return map[int]bool{}
// 	}
// 	visited[startVertex.key] = true
// 	visitCallBack(startVertex.key)

// 	for _, v := range startVertex.adjacent {
// 		if visited[v.key] {
// 			continue
// 		}
// 		visited = DFS(g, v, visitCallBack, visited)
// 	}
// 	return visited
// }

var visitedOrder = [][]int{}

func CallBack(i []int) {
	visitedOrder = append(visitedOrder, i)
}

func DFS(g *Graph, startVertex *Vertex, visitCallBack func([]int), visited map[int]bool, order []int) map[int]bool {

	fmt.Println(visited, startVertex.key)

	if startVertex == nil {
		return map[int]bool{}
	}
	if startVertex.key == 0 {
		visitCallBack(order)
		return map[int]bool{}
	}
	visited[startVertex.key] = true
	order = append(order, startVertex.key)

	for _, v := range startVertex.adjacent {
		if visited[v.key] {
			continue
		}
		visited = DFS(g, v, visitCallBack, visited, order)
	}
	return visited
}

func StartFunctions() {
	ReadFile()
	DataForVetices()
	DataForLinks()

	test.Print()

	visited := map[int]bool{}
	order := []int{}
	fmt.Println(test.vertices[1].key, test.vertices[0].adjacent[0].key)
	visited = DFS(test, test.vertices[0], CallBack, visited, order)
	fmt.Println(visited)
	fmt.Println(visitedOrder)

}
