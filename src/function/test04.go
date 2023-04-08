package main

import "fmt"

func compute(x, y float64, f func(float64, float64) float64) float64 {
	return f(x, y)
}

func f4() {

	x := 3.0
	y := 4.0

	hypot := func(x, y float64) float64 {
		return x * y
	}
	fmt.Println(hypot(x, y)) // 12

	fmt.Println(compute(x, y, hypot)) // 12
	fmt.Println(compute(x, y, func(f float64, f2 float64) float64 {
		return f + f2
	})) // 7
}

//func main() {
//	f4()
//}
