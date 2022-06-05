package main

import (
	// "encoding/json"
	"fmt"
)

type Rectangle struct {
	Length int
	Width int
}

// value receiver: useful when not needing to change values
func (r Rectangle) getArea() int {
	return r.Length * r.Width
}

// pointer receiver: useful when needing to change the values
func (r *Rectangle) setLength(len int) {
	r.Length = len
}

type Person struct {
	// json:"var_name" allows access within the JSON
	// in this case Person.name and Person.Age
	Name string `json:"name"`
	Age uint8
	// ...
}

type Employee struct {
	ID uint64
	Person Person
	Salary uint
	// ...
}

func main() {

	var rect1 Rectangle
	rect1.Length = 5
	rect1.Width = 4

	fmt.Printf("area of rect1: %d\n", rect1.getArea())
	rect1.setLength(7)
	fmt.Printf("New length? %d\n", rect1.Length)
	fmt.Printf("area of rect1: %d\n", rect1.getArea())
	
	// rect2 := Rectangle{}
	// rect2.Length = 6
	// rect2.Width = 9

	// rect3 := Rectangle{5,4}
	// fmt.Printf("rect3 length: %v width: %v\n", rect3.Length, rect3.Width)

	// rect4 := Rectangle{
	// 	Width: 5,
	// 	Length: 4,
	// }
	// fmt.Printf("rect4 length: %v width: %v\n", rect4.Length, rect4.Width)

	// var rect5 = new(Rectangle)
	// rect5.Length = 5

	// var rect6 = &Rectangle{}
	// rect6.Width = 6

	// bob := Employee{
	// 	ID: 78345,
	// 	Salary: 35000,
	// 	Person: Person{
	// 		Name: "Bob",
	// 		Age: 30,
	// 	},
	// }

	// fmt.Printf("Bob: %v\n", bob)

	// susan := Person{
	// 	Name: "Susan",
	// 	Age: 24,
	// }

	// data, err := json.Marshal(susan)
	// if err != nil {
	// 	// error
	// 	fmt.Printf("Error: %v\n", err.Error())
	// }

	// %v is for value in the default format
	// %s is for uninterpreted bytes of slice or string 
	// fmt.Printf("JSON: %v\n", data)
	// fmt.Printf("JSON String: %s\n", data)


}	
