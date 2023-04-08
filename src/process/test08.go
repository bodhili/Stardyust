package main

import (
	"fmt"
	"time"
)

func sw2() {
	fmt.Println("What's Saturday ?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func sw3() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func loop() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	// defer 逆序执行 （LIFO）
	defer func(t time.Time) {
		fmt.Printf("main function is ending, and now month time is: %v\n", t.Month())
	}(time.Now())

	defer fmt.Println("bye bye!")

	sw2()
	sw3()
	loop()
}
