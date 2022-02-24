package main

import "fmt"

func main() {
	section := Section{"cs", "101", "1", []string{}}

	fmt.Println("Simulates basics of class registration")

	var name string

	quit := "no"
	for quit == "no" {
		fmt.Printf("Enter name to enroll in %v: ", section)
		fmt.Scanf("%s ", &name)
		result, err := section.register(name)

		if !result {
			fmt.Println(err.Error())
		}

		fmt.Print("Do you want to quit? (yes/no): ")
		fmt.Scanf("%s ", &quit)
	}
	fmt.Println("The final list of students is:\n", section.Students())
}
