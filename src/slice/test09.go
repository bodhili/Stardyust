package main

import "fmt"

func f9(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// s 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。返回的切片会指向这个新分配的数组
func f91() {
	var s []int
	f9(s)

	// 添加一个空切片
	s = append(s, 0)
	f9(s)

	// 这个切片会按需增长
	s = append(s, 1)
	f9(s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4)
	f9(s)

	// 删除元素
	s = append(s[0:1], s[2:]...)
	f9(s)
}

func forRange(p []int) {
	if p == nil {
		return
	}

	for i, v := range p {
		fmt.Printf("index: %v, value:%v \n", i, v)
	}

	for _, v := range p {
		fmt.Printf("index: %v, value:%v \n", -1, v)
	}

	for i, _ := range p {
		fmt.Printf("index: %v, value:%v \n", i, -1)
	}
}

func main() {
	f91()
	forRange([]int{1, 2, 3})
}
