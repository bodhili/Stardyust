package main

import "fmt"

func p1() {
	// pointer 保存的是值的内存地址 *
	var i *int // 指针定义就分配了内存地址，初始值为 nil
	fmt.Println(i)
	fmt.Printf("i[type]: %T \n", i)
	fmt.Printf("i[p]: %p \n", &i)
	fmt.Printf("i[v]: %v \n", i)
}

//func main() {
//	p1()
//}
