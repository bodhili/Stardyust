package main

import "fmt"

func f5() {
	v := []int{1, 2, 4}

	fmt.Println(len(v))
	fmt.Println(cap(v))
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func f51() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为 0
	s = s[:0] // {2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 拓展其长度
	s = s[:4] //{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 舍弃前两个值
	s = s[2:] //{5, 7, 11, 13}
	printSlice(s)

	s = s[2:] // {11, 13}
	printSlice(s)

	s = s[:2] // {11, 13}
	printSlice(s)

	s = s[2:] //{}

	//	s = s[:4] 切片底层数据被 切为 0 后不可逆 ,, 如果切片引用是空切片，则从底层数组获取数据切片
	printSlice(s)
}

// 切片的长度就是切片自身所包含的元素个数。
// 切片的容量是从切片自身的第一个元素开始，到其底层数组元素末尾的个数
//func main() {
//	f5()
//	f51()
//}
