package main

import "fmt"

type F struct{}

func (f *F) f1() {
	fmt.Println("method")
}

//func main() {
//	f := F{}
//	f.f1()
//}
