package main

import "fmt"


func main() {

	ages := make(map[string]int)
	ages["James"] = 35
	fmt.Printf("James is %d years old\n", ages["James"])

	ages["Susan"] += 1 // Happy Birthday Susan
	fmt.Printf("Susan is %d years old\n", ages["Susan"])

	// addr := &ages["Susan"] // cannot take address of map value

	gpas := map[string]float32 {
		"James": 3.4,
		"Alice": 3.9,
	}
	fmt.Printf("James' GPA is %.2f\n", gpas["James"])
	fmt.Printf("Alice's GPA is %.2f\n", gpas["Alice"])

	var visited map[string]bool
	visited = make(map[string]bool)
	visited["A"] = true
	fmt.Printf("A has been visited? %v\n", visited["A"])

	fruits := []string {
		"bananas",
		"kiwis",
		"apples",
		"strawberries",
		"tomatos",
		"bananas",
	}
	fmt.Printf("Duplicate fruits? %s\n", findDuplicates(fruits))
}

func findDuplicates(words []string) string {
	dupMap := make(map[string]bool)

	for i := 0; i < len(words); i++ {

		if !dupMap[words[i]] {
			dupMap[words[i]] = true
		} else {
			return words[i]
		}
	}
	return ("")
}