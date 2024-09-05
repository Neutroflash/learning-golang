package main

import (
	"fmt"
	"math"
)

func main() {
	//Declaración de variables
	helloMessage := "Hello"
	worldMessage := "World"

	//Println
	fmt.Println(helloMessage, worldMessage)
	
	//Printf
	nombre := "Platzi"
	cursos := 500

	fmt.Printf("%s tiene mas de %d cursos\n ", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n ", nombre, cursos)

	//Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	//Tipo Datos
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

	//Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	//Declaración de variables enteras
	base := 12
	var altura int = 14
	var area int 

	fmt.Println(base, altura, area)

	//Zero Values
	var a int 
	var b float64
	var c string
	var d bool

	fmt.Println(a,b,c,d)

	//Area Cuadrado
	const baseCuadrado = 10
	areaCuadrado := math.Pow(baseCuadrado, 2)
	fmt.Println(areaCuadrado)

	x := 10
	y := 50
	var radio float64 = 3

	//Suma
	result := x + y
	fmt.Println("Suma: ", result)

	//Resta
	result = y - x
	fmt.Println("Resta: ", result)

	//Multiplicación
	result = x * y
	fmt.Println("Multiplicación: ", result)

	//División
	result = y / x
	fmt.Println("División: ", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo: ", result)

	//Incremental
	x++
	fmt.Println("Incremental: ", x)

	//Decremental
	x--
	fmt.Println("Decremental: ", x)

	//Retos
	areaRectangulo := x * y
	areaTrapecio := (x + y) * altura / 2
	areaCirculo :=  pi * math.Pow((radio), 2)

	fmt.Println(areaRectangulo, areaTrapecio, areaCirculo)
}