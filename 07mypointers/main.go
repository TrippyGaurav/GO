package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointers")

	// var ptr *int
	// fmt.Println(ptr)

	num := 23
	var ptr = &num // store reference of num
	fmt.Println("Value of actual pointer is : ", ptr)
	fmt.Println("Value of actual pointer is : ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New Value is : ", num)

}
