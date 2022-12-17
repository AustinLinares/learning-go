package main

import "fmt"

func hello(s string) string {
	return "Hola " + s
}

func plus(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(hello("Austin"))
	fmt.Println(plus(7, 5))
}
