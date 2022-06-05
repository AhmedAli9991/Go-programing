package main

import (
	// "errors"
	"fmt"
	"io/ioutil"
	// "log"
	// "os"
	"strconv"
)

type CustomError struct {
	Message string
	Code int
}

func (c CustomError) Error() string {
	return c.Message + " " + strconv.Itoa(c.Code)
}

func Divide(x, y float64) (float64, error) {
	if y == 0 {
		// return float64(0), errors.New("cannot divide by zero")
		return 0.0, CustomError{"cannot divide by zero", -1}
	} else {
		return x / y, nil
	}
}

func main() {

	file, fileErr := ioutil.ReadFile("file")
	if fileErr != nil {
		fmt.Println("Error reading file:", fileErr)
	} else {
		fmt.Println(string(file))
	}

	value, _ := Divide(7, 0)
	// if divErr != nil {
	// 	// fmt.Println(divErr)
	// 	// os.Exit(-1)
	// 	log.Fatal(divErr)
	// } else {
	// 	fmt.Printf("%.2f\n", value)
	// }
	fmt.Printf("value: %.2f\n", value)
}