package otherpackage

// Animal acceso Publico
type Animal struct {
	Name string
	Year int
	Food string
	Type string
}

func (myAnimal *Animal) DuplicateYear() {
	myAnimal.Year = myAnimal.Year * 2
}
