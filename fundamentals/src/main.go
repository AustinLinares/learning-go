package main

import (
	"fmt"
	"reflect"
	"strconv"
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

func convertData() {
	// string to bool
	var mayorDeEdad string = "true"
	// ignoring error
	boolVal, _ := strconv.ParseBool(mayorDeEdad)
	fmt.Println(boolVal, reflect.TypeOf(boolVal))

	// bool to string
	menorDeEdad := false
	strVal := strconv.FormatBool(menorDeEdad)
	fmt.Println(strVal, reflect.TypeOf(strVal))
}

func declaring(favorite bool) {
	// all variables are declared correctly
	var favoritePokemon string = "Gengar"
	firstPokemon := "Charmander"
	name, firstPokemonGame := "Austin", "oro"
	// var name, firstPokemonGame string = "Austin", "oro"

	var (
		edad   uint8 = 22
		hungry bool  = true
	)

	fmt.Printf("Mi nombre es %v, tengo %v y mi primer juego pokemon fue %v \n", name, edad, firstPokemonGame)

	if hungry {
		fmt.Println("Tengo hambre")
	}

	if favorite {
		fmt.Println("Mi pokemon favorito es", favoritePokemon)
	} else {
		fmt.Println("Mi primer pokemon fue", firstPokemon)
	}
}

func defaultValues() {
	var nombre string
	var edad int
	var peso float64
	var estudiante bool
	fmt.Println("Nombre: ", nombre)         // string vacio
	fmt.Println("Edad: ", edad)             // 0
	fmt.Println("Peso: ", peso)             // 0
	fmt.Println("Estudiante: ", estudiante) //false
}

func scopeTest(param1 string) {
	{
		var var2 = "innerBlock"
		fmt.Println(var2)
	}
	fmt.Println(var1)
	fmt.Println(param1)
}

var var1 = "globalScope"

func pointers() {
	vehiculo1 := "rojo"
	fmt.Println("El vehiculo1 es", vehiculo1) // El vehiculo1 es rojo

	vehiculo2 := vehiculo1
	fmt.Println("El vehiculo2 es", vehiculo2) // El vehiculo2 es rojo

	vehiculo3 := &vehiculo1
	fmt.Println("El vehiculo3 es", *vehiculo3) // El vehiculo3 es rojo

	vehiculo1 = "gris"

	fmt.Println("El vehiculo1 es", vehiculo1, &vehiculo1) // El vehiculo1 es gris 0xc000010230
	fmt.Println("El vehiculo2 es", vehiculo2, &vehiculo2) // El vehiculo2 es rojo 0xc000010250
	fmt.Println("El vehiculo3 es", *vehiculo3, vehiculo3) // El vehiculo3 es gris 0xc000010230
}

func constant() {
	const Pi float32 = 3.1416

	// Pi = 3.14159   // cannot assign to Pi (constant 3.1416 of type float32)

	fmt.Println(Pi)
}

const Pi = 3.1416

func circle(radio float64) (area float64, perimeter float64) {
	area = Pi * radio * radio
	perimeter = 2 * Pi * radio
	return area, perimeter
}

func multiReturn() {
	circleArea, circlePerimeter := circle(8)
	fmt.Println("El area del circulo es: ", circleArea)
	fmt.Println("El perimetro del circulo es: ", circlePerimeter)
}

func variadicPlus(numeros ...int) (total int) {
	// recorrer todos los numeros
	for _, numero := range numeros {
		// en cada iteración sumar al total el valor del numero
		total = numero + total
	}
	// retornar el valor total
	return total
}

func ecoDeLaMontana(mensaje string, iteraciones uint) {
	if iteraciones > 1 {
		ecoDeLaMontana(mensaje, iteraciones-1)
	}
	fmt.Println(mensaje, iteraciones)
}

func lambdas(radio float64) (area func() float64, perimetro func() float64) {

	area = func() float64 {
		return 3.1416 * radio * radio
	}

	perimetro = func() float64 {
		return 2 * 3.1416 * radio
	}

	return
}

func comparisonOperatores() {
	fmt.Println("2 == 2", 2 == 2)               // true
	fmt.Println("2 != 2", 2 != 2)               // false
	fmt.Println("2 < 2", 2 < 2)                 // false
	fmt.Println("2 > 2", 2 > 2)                 //false
	fmt.Println("2 <= 2", 2 <= 2)               // true
	fmt.Println("2 >= 2", 2 >= 2)               // true
	fmt.Println("2 >= 2.50", 2 >= 2.50)         // false
	fmt.Println("2.501 >= 2.50", 2.501 >= 2.50) // true
	fmt.Println("2.50 == 2.5", "2.50" == "2.5") // false
	fmt.Println("2.50 > 2.5", "2.50" > "2.5")   // true
	fmt.Println("2.50 < 2.5", "2.50" < "2.5")   // false
	// fmt.Println("2 >= 2", 2 >= "2") // ERROR:both values should have same data type
}

func aritmeticOperators() {
	fmt.Println("5 + 3", 5+3)   // 8
	fmt.Println("5 - 3", 5-3)   // 2
	fmt.Println("5 * 3", 5*3)   // 15
	fmt.Println("5 / 3", 5/3)   // 1.666
	fmt.Println("5 % 3", 5%3)   // 2
	fmt.Println("10 % 3", 10%3) // 1 // residuo de x entre y
}

func switchEx() {
	var juguete string
	fmt.Println("Elija que tipo de juguete agregar? persona, animal o cosa")
	fmt.Scanln(&juguete)
	switch juguete {
	case "persona":
		fmt.Println("El juguete es una figura de accion")
	case "cosa":
		fmt.Println("El juguete es una cosa")
	case "animal":
		fmt.Println("El juguete es una mascota")
	default:
		fmt.Println("Escoge una opción válida")
	}
}

func main() {
	// var mainVar = "blockVariable"

	// fmt.Println(hello("Austin"))
	// fmt.Println(plus(7, 5))
	// fmt.Println(boolean())
	// input()
	// conditional(false, true)
	// firstArray(0)
	// convertData()
	// declaring(true)
	// defaultValues()
	// scopeTest(mainVar)
	// pointers()
	// constant()
	// multiReturn()
	// fmt.Println(variadicPlus(5, 4, 3))
	// ecoDeLaMontana("Merry Christmas", 3)
	// area, perimetro := lambdas(10)
	// fmt.Println("El area del circulo es", area())
	// fmt.Println("El perimetro del circulo es", perimetro())
	// comparisonOperatores()
	// aritmeticOperators()
	switchEx()
}
