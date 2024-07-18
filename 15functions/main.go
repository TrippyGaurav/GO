package main

import "fmt"

func main() {
	fmt.Println("Introduction to functions")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result : ", result)

	proResult, message := proAdder(1, 2, 4, 5, 6, 8)
	fmt.Println("Pro Result : ", proResult)
	fmt.Println("Pro Message : ", message)
}

func greeter() {
	fmt.Println("Namaste Go Lang")
}

func adder(num1 int, num2 int) int {
	return num1 + num2
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "Hi Pro result function"
}
