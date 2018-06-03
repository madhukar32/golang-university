// Author: Madhukar
// Goes over different way of decalaring variables in golang
package main

import (
	"fmt"
)

func main() {

	//declaring vars by static typing
	var age1 int // simple var declaration "var NAME TYPE"
	fmt.Println("Person 1 age is: ", age1)

	var age2 int = 28 // declaring variable with initial value
	fmt.Println("Person 2 age is: ", age2)

	// type inference
	var age3 = 2
	fmt.Println("Person 3 age is: ", age3)

	// multiple variables can be declared together
	var age4, height4, weight4 = 28, 178, 82
	fmt.Printf("Person 4 age is: %d, height is: %d, weight is: %d \n", age4, height4, weight4)

	// declare vars beloning to different types
	var (
		name          = "golang-university"
		prog_language = "golang"
		age           = 1
	)
	fmt.Printf("Name of repo: %s, lang used: %s, age is: %d \n", name, prog_language, age)

	// short hand declaration (python like)
	awesomness := "go"
	fmt.Println("Awesomness is: ", awesomness)

	// short hand declaration could be used when atleast one of the var is newly declared
	one, two := 1, 2
	fmt.Println("one is:", one, "two is:", two)
	two, three := 2, 3
	fmt.Println("two is:", two, "three is:", three)

}
