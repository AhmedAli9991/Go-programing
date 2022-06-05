package main

import "fmt"


func main() {

	// var arr1 [3]int
	// fmt.Println(arr1)
	// fmt.Println(arr1[0])
	// fmt.Println(arr1[len(arr1) - 1])
	// // fmt.Println(arr1[3]) out of bounds!

	// var arr2 [3]int = [3]int { 1, 2, 3 }
	// fmt.Println(arr2)

	// arr3 := [4]int { 4, 5, 6, 7 }
	// fmt.Println(arr3)
	// // arr3 = arr2 NOT compatible types!

	// arr4 := [...]int {5: -1, 1, 0, 1}
	// fmt.Println(arr4)

	// fmt.Printf("len %d cap %d\n", len(arr4), cap(arr4))

	// months := [...]string {
	// 	"Jan", "Feb", "Mar", "Apr", "May", "Jun",
	// 	"Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
	// }

	// summer := months[5:8]
	// fmt.Printf("Summer: %v\n", summer)
	// q2 := months[3:6]
	// fmt.Println(q2)
	// all := months[:]
	// fmt.Println(all)

	slice := make([]int, 2)
	slice[0] = 1
	slice[1] = 2
	printSlice(slice)
	slice = append(slice, 3)
	printSlice(slice)
	slice = append(slice, 4, 5)
	printSlice(slice)

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Printf("slice[%d] %d\n", i, slice[i])
	// }

	for _, value := range slice {
		fmt.Println(value)
	}

	reverse(slice)
	fmt.Println(slice)

	arr := [...]int { 6, 7, 8, 9, 10}
	reverse(arr[:])
	fmt.Println(arr)
	

}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func printSlice(slice []int) {
	fmt.Printf("Slice %v len %d cap %d\n",
		slice, len(slice), cap(slice))
}