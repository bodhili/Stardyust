package main

import "fmt"

type Audi struct {
	name  string
	price float64
}

func f3() {
	audi := Audi{
		"audi",
		0.0,
	}

	fmt.Println(audi)

	p := &audi
	p.price = 3.1415926

	fmt.Println(*p)
	fmt.Println(audi)
	fmt.Println(audi == *p)
}

//func main() {
//	f3()
//}
