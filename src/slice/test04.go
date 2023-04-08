package main

import "fmt"

func f4() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[:0]
	fmt.Println(s)

	s = s[1:4] //  3, 5, 7  // 变化后 如果s是空切片，丢失空切片，后续变化，还是基于源数据容器进行切片
	fmt.Println(s)

	s = s[:]       //  3, 5, 7
	fmt.Println(s) // 变化后 如果 s 不是空切片，则基于变化后的切片s 进行切片

	s = s[:5] //  3, 5, 7
	fmt.Println(s)

	s = s[1:] // 5
	fmt.Println(s)
}

//func main() {
//	f4()
//}
