package main

import (
	"fmt"
)

type I1 interface {
	M1()
}

type T1 struct {
	S string
}

func (t *T1) M1() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I1) {
	fmt.Printf("(%v, %T)\n", i, i)
}

//func main() {
//	var i I1
//
//	i = &T1{"Hello"}
//	describe(i)
//	i.M1()
//}
