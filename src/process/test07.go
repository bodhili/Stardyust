package main

import (
	"fmt"
	"runtime"
)

func sw() {
	// 变量初始化可在 if switch 第一个分支进行初始化 (;)
	// switch 默认不进行穿透执行，case 匹配自动break
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	//	fallthrough 穿透一次
	case "Linux":
		fmt.Println("Linux.")
		fallthrough
	default:
		fmt.Println("Mac OS.")
	}
}

func sw1(x int) {
	switch x * 2 {
	case 4:
		fmt.Println("3")
	}
}

//func main() {
//	fmt.Println("Go runs on:")
//	sw()
//	sw1(2)
//}
