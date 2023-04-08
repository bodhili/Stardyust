package main

import "fmt"

type test struct {
	X int
	Y int
}

func f1() {
	t1 := test{
		X: 1,
		Y: 2,
	}

	t2 := new(test) //new 结果是: 指针类型, 值: &v
	t3 := &t2       //new 结果是: 指针类型, 值: &v

	fmt.Println(t1)
	fmt.Println(t3)
	fmt.Printf("t2[type]: %T, t2[v]: %v, t2[p]: %v \n", t2, t2, *t2)
}

//func main() {
//	f1()
//}
