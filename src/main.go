package main

import (
	"fmt"
	"math"
)

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func isPar(a int) {
	if a%2 == 0 {
		fmt.Println("Es par")
	} else {
		fmt.Println("Es impar")
	}
}

func isLogin(user, password string) {
	if user == "Miguel" && password == "loli1964" {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

func main() {
	isLogin("Miguel", "loli196d4")
	isPar(2223)
	tripleArgument(1, 2, "Hola")
	normalFunction("Aprendiendo Golang!")
	value := returnValue(2)
	fmt.Println("Value: ", value)

	value1, value2 := doubleReturn(2)
	fmt.Println(value1, value2)

	value3, _ := doubleReturn(3)
	fmt.Println(value3)

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

	fmt.Println(a, b, c, d)

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
	areaCirculo := pi * math.Pow((radio), 2)

	fmt.Println(areaRectangulo, areaTrapecio, areaCirculo)

	//For condicional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	//For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	//For forever
	/*
		counterForever := 0
		for {
			fmt.Println(counterForever)
			counterForever++
		}
	*/

	for i := 10; i >= 0; i-- {
		fmt.Println(i)
	}

	counter2 := 10
	for counter2 >= 0 {
		fmt.Println(counter2)
		counter2--
	}

	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	//and
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	}

	//or
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Claro que si")
	} else {
		fmt.Println("Soluciona el problema")
	}

	//Convertir texto a numero
	/*
		value, err := strconv.Atoi("5sada")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("value:" ,value)
	*/

	/*
		modulo := 4 % 2
		switch modulo {
		case 0:
			fmt.Println("Es Par")
		default:
			fmt.Println("Es Impar")
		}
	*/

	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es Par")
	default:
		fmt.Println("Es Impar")
	}

	//Sin Condición
	valuew := 50
	switch {
	case valuew > 100:
		fmt.Println("Es Mayor a 100")
	case valuew < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No Condición")
	}

	//Defer
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//Continue-Break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		//Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		//Break
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}

	//Array
	var array [4]int 
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	//Slices
	slice := []int{0,1,2,3,4,5,6}
	fmt.Println(slice, len(slice), cap(slice))

	//Métodos en Slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	//Append
	slice = append(slice, 7)
	fmt.Println(slice)

	//Append nueva list
	newSLice := []int{8,9,10}
	slice = append(slice, newSLice...)
	fmt.Println(slice)
}
