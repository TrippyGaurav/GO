package main

import "fmt"

func main() {
	fmt.Println("Introduction to Arrays ")

	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[3] = "Mango"

	fmt.Println("Fruits : ", fruits)
	fmt.Println("Total Fruits: ", len(fruits))

	var vegetables = [3]string{"Potato", "Tomato", "Mushrooms"}
	fmt.Println("Vegetables : ", vegetables)
	fmt.Println("Total Vegetables: ", len(vegetables))

}
