package main

import "fmt"

func main() {
	fmt.Println("Introduction to Loops")

	// days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value id %v\n", index, day)
	// }

	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}
		fmt.Println("Value is : ", rougueValue)
		rougueValue++
	}

lco:
	fmt.Printf("Hello")
}
