package main

import "fmt"

type I2 interface {
	M2()
}

type T2 struct {
	S string
}

func (t *T2) M2() {
	//if t == nil {
	//	fmt.Println("<nil>")
	//	return
	//}   不这样判断，将会NPE

	fmt.Println(t.S)
}

//func main() {
//	var i2 I2  // nil
//	var t2 *T2 // nil
//
//	i2 = t2
//
//	i2.M2()
//}
