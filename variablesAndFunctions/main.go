package main

import (
	"fmt"
)

//package level variable 
var myName string

func main() {

	fmt.Println("Hello World")

	var whatToSay string //string variable decleration
	var i int32  //integer variable decleration(signed int 32 bits)
        
	//assigning values to variable
	whatToSay = "Goodbye, cruel world" 
	i = 12
	//myName ="dingle dong"
        
	//printing the values
        fmt.Println(whatToSay)	
	fmt.Println(i)
	//fmt.Println(myName)

	whatDidCatSay, ageOfTheCat := catSays()
	fmt.Sprintln("Cat says: %s and the age of the cat is %d", whatDidCatSay,ageOfTheCat)
}

//function decleration and defination
//can return more than one value
func catSays() (string, int) {
	return "meow",3
}
