package main

import (
	"fmt"
)

func main() {
	fmt.Println("Introduction to Slices")

	// var fruits = []string{"Apple", "Mango", "Banana"}
	// fmt.Printf("Type of fruits : %T\n", fruits)

	// fruits = append(fruits, "Graphes")
	// fmt.Println(fruits)

	// fruits = append(fruits[:3])
	// fmt.Println(fruits)

	// highScores := make([]int, 4)
	// highScores[0] = 234
	// highScores[1] = 945
	// highScores[2] = 465
	// highScores[3] = 867
	// highScores[4] = 777

	// highScores = append(highScores, 555, 666, 321)
	// fmt.Println("Scores : ", highScores)

	// fmt.Println(sort.IntsAreSorted(highScores))
	// sort.Ints(highScores)
	// fmt.Println(highScores)

	// How to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...);
	fmt.Println(courses)
	
}
