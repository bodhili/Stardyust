package main

import (
	"fmt"
)

func main() {
	t := DeferFunc3(1)
	fmt.Println(t)
	fmt.Println(DeferFunc2(1))

	defer_call()

	fmt.Println("main 正常结束")
}

func _error() {
	fmt.Println("init...")

	recover()

	fmt.Println("end...")
}

func DeferFunc2(i int) (result int) {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

func defer_call() {
	defer func() { fmt.Println("defer: panic 之前1") }()
	defer func() { fmt.Println("defer: panic 之前2") }()

	_error()

	panic("异常内容") //触发defer出栈

	defer func() { fmt.Println("defer: panic 之后，永远执行不到") }()

}
