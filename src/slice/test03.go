package main

import "fmt"

func f3() {
	q := []int{2, 3, 5, 7, 11, 13} // 1. [6]int {2, 3, 5, 7, 11, 13} // 2. 切片 q 对数组的引用，，原理：申明切片的同时会创建对应类型的容器来存储数据, 切片变量持有对创建的容器的引用
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
}

func f31() {
	s := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{3, true},
	}

	fmt.Println(s)
	fmt.Printf("s[type]: %T \n", s)
}

type Vertex struct {
	name string
	age  int
	die  bool
}

func f32() {
	v := []Vertex{
		{"tom", 20, true},
		{"kk", 20, false},
	}

	p := &v

	(*p)[0].age = 200

	fmt.Println(*p)
	fmt.Println(v)
	fmt.Println(len(v))
}

//func main() {
//	f3()
//	f31()
//	f32()
//}
