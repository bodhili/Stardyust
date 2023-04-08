package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func f1() {
	m = make(map[string]Vertex) // make 先分配

	m["tom"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["tom"])
	vertex := m["Bell Labs"]
	fmt.Println(vertex)
}

func f11() {
	m := map[string]Vertex{ // 直接初始化
		"tom": {
			1.0, 2.0,
		},
		"Google": Vertex{
			3.0, 4.0,
		},
	}

	fmt.Println(m)
}

//func main() {
//	f1()
//	f11()
//}
