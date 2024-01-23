<<<<<<< HEAD
package main

import (
	"log"
)

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)

	changeUsingPointer(&myString) //refrence
	log.Println("myString is updated to",myString)
}

func changeUsingPointer(s *string) { //pointer
	log.Println("s is set to :", s)
	newValue := "Red"
	*s = newValue
}
=======
package main

import (
	"log"
)

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)

	changeUsingPointer(&myString)
	log.Println("myString is updated to",myString)
}

func changeUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}
>>>>>>> a784fce (added basics of variables , functions and pointers)
