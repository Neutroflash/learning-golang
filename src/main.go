package main

import (
	pk "curso_golang_platzi/src/mypackage"
	pk2 "curso_golang_platzi/src/otherpackage"
	"fmt"
	"math"
	"strings"
	"sync"
	"time"
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

func isPalindromo(text string) {
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if strings.ToLower(text) == strings.ToLower(textReverse) {
		fmt.Println("Es Palindromo")
	} else {
		fmt.Println("No es Palindormo")
	}
}

type car struct {
	brand string
	year  int
}

type pc struct {
	ram   int
	disk  int
	brand string
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

type figuras2D interface {
	area() float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figuras2D) {
	fmt.Println("Area: ", f.area())
}

func (myPc pc) ping() {
	fmt.Println(myPc.brand, "Pong")
}

func (myPc *pc) duplicateRam() {
	myPc.ram = myPc.ram * 2
}

/*
func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB DISCO y es una %s", myPc.ram, myPc.disk, myPc.String())
}
*/

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func say2(text string, canal chan<- string) {
	canal <- text
}

func message(text string, c chan string) {
	c <- text
}

func main() {
	canal2 := make(chan string, 2)
	canal2 <- "Mensaje 1"
	canal2 <- "Mensaje 2"

	fmt.Println(len(canal2), cap(canal2))

	//Range y Close
	close(canal2)

	//canal2 <- "Mensaje 3"
	for message := range canal2 {
		fmt.Println(message)
	}

	//select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("mensaje1", email1)
	go message("mensaje2", email2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibidio de email2", m2)
		}
	}

	canal := make(chan string, 1)
	fmt.Println("Test Email")

	go say2("Bye", canal)
	fmt.Println(<-canal)

	var wg sync.WaitGroup

	fmt.Println("Hello")
	wg.Add(1)
	go say("World", &wg)
	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("Adios")
	time.Sleep(time.Second * 1)

	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 2, altura: 4}

	calcular(myCuadrado)
	calcular(myRectangulo)

	//Lista interfaces
	myInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterface...)

	var myAnimal pk2.Animal
	myAnimal.Name = "Fido"
	myAnimal.Year = 2
	myAnimal.Food = "Ricocan"
	myAnimal.Type = "Perro"
	fmt.Println(myAnimal)

	myAnimal.DuplicateYear()
	fmt.Println(myAnimal.Year)

	var myCar2 pk.CarPublic
	myCar2.Brand = "Ferrari"
	myCar2.Year = 2020
	fmt.Println(myCar2)

	pk.PrintMessage("Hola Platzi")

	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	//Otra Manera
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)

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
	slice := []int{0, 1, 2, 3, 4, 5, 6}
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
	newSLice := []int{8, 9, 10}
	slice = append(slice, newSLice...)
	fmt.Println(slice)

	slice2 := []string{"hola", "que", "hace"}
	for i := range slice2 {
		fmt.Println(i)
	}

	//ama
	//amor a roma
	isPalindromo("Amor a Roma")

	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	//Recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar un valor
	val, ok := m["Jose"]
	fmt.Println(val, ok)

	//Punteros
	val1 := 50
	val2 := &val1

	fmt.Println(val2)
	fmt.Println(*val2)

	*val2 = 100
	fmt.Println(val1)

	myPc := pc{ram: 16, disk: 200, brand: "MSI"}
	fmt.Println(myPc)

	myPc.ping()

	fmt.Println(myPc)
	myPc.duplicateRam()
	fmt.Println(myPc)
	myPc.duplicateRam()
	fmt.Println(myPc)
}
