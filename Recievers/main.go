package main

import (
	"log"
)

type myStruct struct {
	FirstName string
}

//defining a reciever
func (m* myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct 
	myVar.FirstName = "Dingle"

	myVar2 := myStruct {
		FirstName: "Donger",
	}

	log.Println("myVar is set to",myVar.printFirstName())
	log.Println("myVar2 is set to",myVar2.printFirstName())
}

