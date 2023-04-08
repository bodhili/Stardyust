package main

import "fmt"

func f1() {
	a1 := [...]int{1, 2, 3, 4, 5, 5, 6}

	// 包含下边界值，但是不包含上边界值
	s1 := a1[:]
	s2 := a1[1:]
	s3 := a1[:4]
	s4 := a1[1:5]

	fmt.Println(a1)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}

//func main() {
//	f1()
//}
