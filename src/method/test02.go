package main

import "fmt"

type F1 struct {
	name string
	age  int
}

func (f *F1) f2() string {
	f.age = f.age + 10
	return fmt.Sprintf("name: %v, age:%d \n", f.name, f.age)
}

func (f *F1) sum() {
	f.age = f.age * 10
}

func (f F1) sum0() {
	f.age = f.age * 10
}

//func main() {
//	f1 := F1{
//		"tom", 20,
//	}
//
//	fmt.Printf("f1[type]: %T \n", f1)
//
//	ret := f1.f2()
//	fmt.Println(ret)
//
//	//f1.sum()
//	//fmt.Println(f1)
//
//	f1.sum0()
//	fmt.Println(f1)
//
//	ret0 := f1.f2()
//	fmt.Println(ret0)
//}
