package main

import "fmt"

func hello(s string) string {
	return "Hola " + s
}

func plus(a int, b int) int {
	return a + b
}

func boolean() bool {
	var b bool
	return b
}

func conditional(a bool, b bool) {
	if a {
		fmt.Println("A is true")
	} else if b {
		fmt.Println("B is true")
	} else {
		fmt.Println("None is true")
	}
}

func input() {
	var peso float32
	var nombre, apellido string
	fmt.Println("Ingresa tu nombre")
	fmt.Scanf("%s", &nombre)
	fmt.Println("Ingresa tu apellido")
	fmt.Scanf("%s", &apellido)
	fmt.Println("Bienvenido", nombre, apellido)

	fmt.Println("Ingresa tu peso")
	fmt.Scanf("%f", &peso)
	fmt.Println("El peso es:", peso, "kg")
}

func main() {
	// fmt.Println(hello("Austin"))
	// fmt.Println(plus(7, 5))
	// fmt.Println(boolean())
	// input()
	conditional(false, true)
}
