package main

import "fmt"

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key int
	adjacent []*Vertex
}

func (g *Graph) addVertex(value int) {

	if !contains(g.vertices,value) {
		g.vertices = append(g.vertices, &Vertex{key:value})
	}
	
}

func (g *Graph) addEdge(from, to int) {

	fromVertex := getVertex(g.vertices,from)
	toVertex := getVertex(g.vertices,to)

	if fromVertex == nil || toVertex == nil {
		return
	}

	if !contains(fromVertex.adjacent,to) {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}

}

func getVertex(s []*Vertex,k int) *Vertex {

	for _,val := range s {
		if val.key == k {
			return val
		}
	}
	return nil
}

func contains(s []*Vertex,value int) bool {

	for _,val := range s {
		if val.key == value {
			return true
		}
	}
	return false
}


func (g *Graph) Print() {

	for _,val := range g.vertices {
		fmt.Printf("%v :",val.key)
		for _,val := range val.adjacent {
			fmt.Printf(" %v ",val.key)
		}
		fmt.Println()
	}
}

func main() {

	g := Graph{}
	g.addVertex(1)
	g.addVertex(2)
	g.addVertex(3)
	g.addVertex(4)
	g.addVertex(5)

	g.addEdge(1,2)
	g.addEdge(1,3)
	g.addEdge(1,4)
	g.addEdge(2,4)
	g.addEdge(2,1)
	g.addEdge(2,3)
	g.addEdge(4,5)
	g.addEdge(3,4)

	

	// g.Print()

	g.BFS(2)

	g.DFS(2)
}

type queue []int

func (g *Graph) BFS(start int) {

	q := queue{}
	q = append(q, start)
	visited := make(map[int]bool)
	visited[start] = true

	for len(q) != 0 {

		vertex :=	q[0]
		q = q[1:]

		fmt.Print(vertex," ")

		for _,neighbors := range getVertex(g.vertices,vertex).adjacent {
			if !visited[neighbors.key] {
				visited[neighbors.key] = true
				q = append(q, neighbors.key)
			}
		}
	}
	fmt.Println()

}


func (g *Graph) DFS(start int) {

	s := []int{}

	s = append(s, start)
	visited := make(map[int]bool)
	visited[start] = true

	for len(s) != 0 {

		vertex := s[len(s)-1]
		s = s[:len(s)-1]
		fmt.Print(vertex," ")

		for	_,neighbors := range getVertex(g.vertices,vertex).adjacent {
			if !visited[neighbors.key] {
				visited[neighbors.key] = true
				s = append(s, neighbors.key)
			}
		}
	}
	fmt.Println()
	
}

