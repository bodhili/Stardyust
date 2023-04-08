package main

import "fmt"

func f2() {
	m := make(map[string]int)

	m["tom"] = 1
	fmt.Println(m["tom"])

	m["tom"] = 4
	fmt.Println(m["tom"])

	delete(m, "tom")
	fmt.Println(m["tom"])

	v, ok := m["tom"] // 双赋值判断 key 是否存在，OK: true / false
	if ok {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
		fmt.Println(v, ok)
	}

	m1 := map[string]string{}

	v1, ok1 := m1["joy"]
	if ok1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
		fmt.Println(v1 == "") // true
		fmt.Println(len(v1))  // 0
		fmt.Println(&v1)      // 0x1400010a230
		fmt.Println(v1, ok1)
	}
}

func main() {
	f2()
}
