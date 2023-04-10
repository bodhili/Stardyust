package main

type f interface {
	sum() int
}

type U struct{}

type MyInt int

func (f MyInt) sum() int {
	return -1
}

func (u *U) sum() int {
	return -3
}

//func (u U) sum() int {
//	return -2
//}

//func main() {
//	var _face f
//
//	m := MyInt(-1)
//	sum := m.sum()
//
//	println(sum)
//
//	u := &U{}
//	_face = u
//	println(_face.sum())
//}
