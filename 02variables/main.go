package main

import (
	"fmt"

)

func main() {
	// for strings
	var username string = "prajwal"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	// small value
	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("Variable if of type: %T \n", smallVal)

	// small float
	var smallfloat float32 = 255.455
	fmt.Println(smallfloat)
	fmt.Printf("Variable if of type: %T \n", smallVal)


	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type:%T \n",anotherVariable)

	// implicit type
	var website="www.google.com"
	fmt.Printf(website)

	// no var style
	numberOfUser:=300000.00
	fmt.Println(numberOfUser)

}
