package main

import (
	"fmt"
	"reflect"
)

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

func firstArray(i uint8) {
	var colores [3]string
	colores[0] = "Negro"
	colores[1] = "Azul"
	colores[2] = "Gris"

	fmt.Println("El color elegido fue", colores[i], reflect.TypeOf(colores[i]))
	fmt.Println("El último color es", colores[len(colores)-1])

	fmt.Printf("El primer color es %v y es una variable de tipo %T \n", colores[0], colores[0])
}

func main() {
	// fmt.Println(hello("Austin"))
	// fmt.Println(plus(7, 5))
	// fmt.Println(boolean())
	// input()
	// conditional(false, true)
	firstArray(0)
}
