package main

import "fmt"

func main() {
	fmt.Println("hello word...")
}

func Test(name string) string {
	return test(name)
}

func test(name string) string {
	if name == "" || len(name) == 0 {
		return ""
	}

	return name
}
