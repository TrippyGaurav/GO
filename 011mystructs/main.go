package main

import "fmt"

func main() {
	fmt.Println("Strcuts in golang")
	// No inheritance in golang; No super or parent

	gaurav := User{"Gaurav", "gaurav@gmail.com", true, 23}
	fmt.Println(gaurav)
	fmt.Printf("Gaurav details are : %+v\n", gaurav)
	fmt.Printf("Name : %v\n", gaurav.Name)
	fmt.Printf("Email : %v\n", gaurav.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
