package main

import "fmt"

type Vertex struct {
	x, y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{x: 1}
	v3 = Vertex{y: 2}
	v4 = Vertex{}

	p = &Vertex{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3, v4, *p)
}
