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

// func DFS(g *Graph, startVertex *Vertex, visitCallBack func([]int), CheckVisited map[int]bool, order []int) map[int]bool {

// 	fmt.Println(CheckVisited, startVertex.key)

// 	if startVertex == nil {
// 		return map[int]bool{}
// 	}

// 	if startVertex.key == 0 {
// 		visitCallBack(order)
// 		fmt.Println("-------NEw PATH-------")
// 		return CheckVisited
// 		// return map[int]bool{}
// 	}

// 	CheckVisited[startVertex.key] = true
// 	fmt.Println("VISITED")
// 	order = append(order, startVertex.key)

// 	for _, v := range startVertex.adjacent {
// 		if CheckVisited[v.key] {
// 			continue
// 		}
// 		CheckVisited = DFS(g, v, visitCallBack, CheckVisited, order)
// 		fmt.Println(CheckVisited, v.key, order)
// 	}
// 	return CheckVisited
// }

func ContainsVertice(vertice int, order []int) bool {
	for _, val := range order {
		if vertice == val {
			return true
		}
	}
	return false
}

var visitedOrder = [][]int{}

func CallBack(i []int) {
	visitedOrder = append(visitedOrder, i)
}

func RewriteSlice(i []int) []int {
	neworder := []int{}
	for _, val := range i {
		neworder = append(neworder, val)
	}
	return neworder
}

func RewriteOrder(i []int) {
	shortest = nil
	for _, val := range i {
		shortest = append(shortest, val)
	}
}

func DFS(g *Graph, startVertex *Vertex, visitCallBack func([]int), order []int) {
	fmt.Println(startVertex.key, order)
	fmt.Println("Order:", visitedOrder)

	if startVertex == nil {
		fmt.Println("NIL value!")
		return
	}

	if startVertex.key == 0 {
		order = append(order, startVertex.key)
		neworder := RewriteSlice(order)
		visitCallBack(neworder)
		return
	}
	order = append(order, startVertex.key)

	for _, v := range startVertex.adjacent {
		if ContainsVertice(v.key, order) {
			continue
		}
		DFS(g, v, visitCallBack, order)
	}
	return
}

func CheckForDuplicateRooms(paths [][]int, newpath []int) bool {
	if len(paths) == 0 {
		return true
	}
	for _, val := range paths {
		for _, room := range val[1 : len(val)-1] {
			for _, newroom := range newpath[1 : len(newpath)-1] {
				if room == newroom {
					return false
				}
			}
		}
	}
	return true
}

var shortest = []int{}

func SearchForValidPaths(orders [][]int, startVertex *Vertex) [][]int {
	paths := [][]int{}
	for i := 0; i < len(startVertex.adjacent); i++ {
		for index, order := range orders {
			if order[1] == startVertex.adjacent[i].key {
				if CheckForDuplicateRooms(paths, order) {
					fmt.Println(order[1], startVertex.adjacent[i].key, "--------")
					if len(shortest) != 0 && (len(shortest) > len(orders[index])) {
						RewriteOrder(order)
					} else if len(shortest) == 0 {
						RewriteOrder(order)
					}
				}
			} else {
				continue
			}
		}
		paths = append(paths, shortest)
		shortest = nil

	}
	return paths
}

func StartFunctions() {
	ReadFile()
	DataForVetices()
	DataForLinks()

	test.Print()

	CheckVisited := map[int]bool{}
	order := []int{}
	fmt.Println(test.vertices[1].key, test.vertices[0].adjacent[0].key)
	// CheckVisited = DFS(test, test.vertices[0], CallBack, CheckVisited, order)
	DFS(test, test.vertices[0], CallBack, order)
	result := SearchForValidPaths(visitedOrder, test.vertices[0])

	fmt.Println(CheckVisited)
	fmt.Println("Order:")
	for index, val := range visitedOrder {
		fmt.Print(index + 1)
		fmt.Println(val)
	}

	fmt.Println(result)

}
