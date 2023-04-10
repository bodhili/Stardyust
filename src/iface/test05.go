package main

import "fmt"

// Null interface 空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）
type Null interface{}

func describe0(n Null) {
	fmt.Printf("(%v, %T)\n", n, n)
}

//func main() {
//	var i interface{}
//	describe0(i)
//
//	i = 42
//	describe0(i)
//
//	i = "hello"
//	describe0(i)
//}
