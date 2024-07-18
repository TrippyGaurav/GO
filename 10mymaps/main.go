package main

import "fmt"

func main() {
	fmt.Println("Introduction to Maps")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	// fmt.Println("languages : ", languages)
	// fmt.Println("JS : ", languages["JS"])

	// delete(languages, "RB")
	// fmt.Println("languages : ", languages)

	for key, value := range languages{
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	for _, value := range languages{
		fmt.Printf("For key value is %v\n", value)
	}

}
