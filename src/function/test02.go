package main

import (
	"errors"
)

func sum(x, y int) (int, error) {
	if x < 0 || y < 0 {
		return -1, errors.New("cal ret is error")
	}

	return x + y, nil
}

func swap(x, y int) (int, int) {
	x = x ^ y // x = x ^ y
	y = x ^ y // y = x ^ y ^ y
	x = x ^ y // x = x ^ y ^ (x ^ y ^ y)   y

	return x, y
}

func Swap(x, y string) (string, string) {
	return y, x
}

//func main() {
//	s1, s2 := Swap("a", "b")
//	fmt.Printf("s1: %v \n", s1)
//	fmt.Printf("s2: %v \n", s2)
//	fmt.Println(swap(1, 2))
//}
