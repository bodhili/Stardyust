package main

import "fmt"

type User struct {
	name string
	age  int
	die  bool
}

func f2() {
	var user = User{
		"tom",
		20,
		true,
	}

	user.die = false

	fmt.Println(user)
}

//func main() {
//	f2()
//}
