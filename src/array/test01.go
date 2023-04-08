package main

import "fmt"

func f1() {
	var a [10]int
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T \n", a)
	fmt.Printf("%p \n", &a)

	p := &a
	p[9] = 9

	fmt.Println(a)
}

func f2() {
	var a [2]string
	a[0] = "hello"
	a[1] = "word"

	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Println(a[1])

	primes := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(primes)

	var p2 [2]int
	primes1 := [...]int{1, 2}

	//	p2 = primes1  error , type not match

	p2[0] = 1
	p2[1] = 2

	fmt.Println(len(p2))
	fmt.Println(primes1)
}

func f3() {
	a1 := [...]int{1, 2, 3}
	for i := 0; i < len(a1); i++ {
		fmt.Println(a1[i])
	}
}

func main() {
	f1()
	f2()
	f3()
}
