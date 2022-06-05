package main

import (
	"fmt"
	"strconv"
)

func main() {
	// n := 5

	// for i := 0; i < n; i++ {
	// 	fmt.Printf("Hello.\n")
	// }

	
	// i := 0
	// for i < n {
	// 	fmt.Println("Hello.")
	// 	i++ // i = i + 1
	// }
	// fmt.Printf("the value of i is %d\n", i)
	/*
	var number, reverse int
	fmt.Printf("Enter a Number to reverse: ")
	fmt.Scan(&number)
	// assume number is 21
	for number != 0 {
		// 1st iteration : reverse is 0
		// 2nd iteration : 1 X 10 = 10
		reverse = reverse * 10
		// 1st iteration : 0 + (21 % 10) 1 = reverse is 1
		// 2nd iteration : 10 + (2 % 10) 2 = reverse is 12
		reverse = reverse + number % 10
		// 1st iteration : 21 / 10 = 2
		// 2nd iteration : 2 / 10 = 0
		number = number / 10
	}

	fmt.Printf("Reverse Number: %d\n", reverse)
	*/

	var input string
	fmt.Printf("Enter a number\n")
	fmt.Scan(&input)
	num, err := strconv.Atoi(input)
	if err != nil {
		for {
			fmt.Printf("Invalid input, please enter a number: \n")
			fmt.Scan(&input)
			num, err = strconv.Atoi(input)
			if err != nil {
				continue
			}
			if err == nil {
				break
			}
		}
	}
	fmt.Printf("You entered: %d\n", num)

	
	/*
	nums := []int { 1, 2, 3}
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	for _, num := range nums {
		fmt.Printf("%d\n", num)
	}
	*/
}