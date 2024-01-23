package main

import (
	"log"
)

//variable shadowing
var s string = "seven" //global/package level variable

func main() {
        var s2 string = "six"

	log.Println("s is",s)
	log.Println("s2 is",s2)

	s:= "eight"//local variable of main function
        log.Println("s now is",s)

	saySomething("hueueu")
	saySomething2("heheh")
}

func saySomething(s string) (string, string) {
	log.Println("s from the saySomething func is ",s) //will use the parameter which is passed instead of package level variable
	return s, "world"
}

func saySomething2(s3 string) (string, string) {
	log.Println("s from the saySomething func is ",s) //will use the global/package level variale
	return s,"world"
}
