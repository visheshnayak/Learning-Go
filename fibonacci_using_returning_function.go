package main

import (
	"fmt"
)

func fibo() func(int) int {
	a := 0
	b := 1

	return func(iter int) int {
		switch iter {
		case 0:
			return a
		case 1:
			return b
		default:
			c := a + b
			a = b
			b = c
			return c
		}
	}
}

func main() {
	f := fibo()

	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
