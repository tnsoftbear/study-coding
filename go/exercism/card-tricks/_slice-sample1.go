package main

import "fmt"

func main() {
	customers := [4]string{"John Doe", "Helmuth Verein", "Dany Beril", "Oliver Lump"}
	customersSlice := customers[0:1]
	fmt.Println(customersSlice)
	// modify original array
	customers[0] = "John Doe Modified"
	fmt.Println("After modification of original array")
	fmt.Println(customersSlice)
	fmt.Println(customers) // 0 indexed value is modified

	hotelName := "Go Dev Hotel"
	s := hotelName[0:6] // s is immutable
	fmt.Println(s)
	hotelName = "Java Dev Hotel"
	fmt.Println(s) // s is not modified
	fmt.Println(hotelName)
}
