package main

import "fmt"

// 切片就像数组的引用
// 切片并不存储任何数据，它只是描述了底层数组中的一段
// 更改切片的元素会修改其底层数组中对应的元素
// 与它共享底层数组的切片都会观测到这些修改

func f2() {
	a1 := [...]int{1, 2, 3} // 等价 a1 := []int {1, 2, 3}

	s1 := a1[:2]
	s1[0] = 100

	fmt.Println(len(a1))
	fmt.Println(len(s1))
	fmt.Println(s1)
	fmt.Println(a1)

	fmt.Printf("s1[type]: %T \n", s1)
	fmt.Printf("a1[type]: %T \n", a1)
	fmt.Printf("a1[p]: %p \n", &a1[0])
	fmt.Printf("a1[v]: %v \n", *&a1[0])

	fmt.Println(a1[0] == *&a1[0])
}

func f21() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2] // "John", "Paul",
	b := names[1:3] // "Paul", "George",
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	fmt.Println(&names[0])
	fmt.Println(&a[0])
	fmt.Println(&b[0])

	fmt.Println(&b[0] == &a[1]) // true, 说明切片只是持有源数据的内存地址的引用，并没有copy副本单独保存值
}

//func main() {
//	f2()
//	f21()
//}
