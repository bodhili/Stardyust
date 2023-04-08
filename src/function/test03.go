package main

func split(sum int) (x, y int) {
	x = sum << 2 / 9
	y = sum - x
	return
}

//func main() {
//	println(split(8))
//}
