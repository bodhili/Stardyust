package main

import "fmt"

func f7(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func f71() {
	a := make([]int, 5) // len:5 cap:5
	f7("a", a)

	b := make([]int, 0, 5) // len:0 cap:5
	f7("b", b)

	c := b[:2]
	f7("c", c) // len:2 cap:5

	d := c[2:5] //// len:3 cap:3
	f7("d", d)
}

//func main() {
//	f71()
//}
