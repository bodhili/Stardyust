package main

import "fmt"

func p2() {
	i := 10
	p := &i // 内存地址

	var x *int      // 指针类型存储内存地址
	fmt.Println(&x) // 0x14000122020
	fmt.Println(&p) // 0x14000122018

	x = &i // 赋值内存地址

	// 内存地址
	fmt.Println(&i) // 0x1400012c008 对应的值是 10
	fmt.Println(p)  // 0x14000122020 对应的内存空间存储的是 i的内存地址值: 0x1400012c008
	fmt.Println(x)  // 0x14000122018 对应的内存空间存储的是 i的内存地址值: 0x1400012c008

	fmt.Printf("p[type]: %T \n", p)
	fmt.Printf("p[v]: %v \n", *p) //从内存地址取值

	*p = 20 // 给对应的指针内存地址赋值

	fmt.Println(i)
	fmt.Println(*p)
	fmt.Println(*x)
}

//func main() {
//	p2()
//}
