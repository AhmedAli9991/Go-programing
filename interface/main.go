package main

import "fmt"

type printer interface {
	printInfo()
}

type Book struct {
	Title string
	Price float32
}

type Drink struct {
	Name string
	Price float32
}

type Toy struct {
	Name string
	Price float32
}

func (d Drink) printInfo() {
	fmt.Printf("Drink: %s Price: %.2f\n", d.Name, d.Price)
}

func (b Book) printInfo() {
	fmt.Printf("Title: %s Price: %.2f\n", b.Title, b.Price)
}

func (t Toy) printInfo() {
	fmt.Printf("Toy: %s Price: %.2f\n", t.Name, t.Price)
}

func empty(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I'm a string: %s\n", i)
	case int:
		fmt.Printf("I'm an int %d\n", i)
	case Book:
		// Type Assertion
		// i.(T) -> we are asserting that i contains the type T
		fmt.Printf("I'm a book: %s\n", i.(Book).Title)
	}
	
}

func main() {

	gunslinger := Book {
		Title: "The Gunslinger",
		Price : 4.75,
	}

	coffee := Drink {
		Name: "Coffee",
		Price: 2.50,
	}

	rubixCube := Toy {
		Name: "Rubix Cube",
		Price: 5.00,
	}

	gunslinger.printInfo()
	coffee.printInfo()
	// rubixCube.printToy()

	info := []printer { gunslinger, coffee, rubixCube }

	info[0].printInfo()
	info[1].printInfo()
	info[2].printInfo()

	empty("James")
	empty(23)
	empty(gunslinger)

}