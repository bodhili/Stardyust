package main

import "fmt"

func f6() {
	var s []int
	fmt.Printf("s[p]: %p \n", &s)

	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("slice is nil...")
	}
}

//func main() {
//	f6()
//}
