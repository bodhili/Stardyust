package main

import (
	"errors"
)

func add(x, y int) int {
	return x + y
}

func Add(x, y int) (int, error) {
	if x < 0 || y < 0 {
		return -1, errors.New("cal ret is error")
	}

	return x * y, nil
}

//func main() {
//	ret := add(1, 2)
//	fmt.Println(ret)
//}
