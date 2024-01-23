package main

import (
	"log"
	"fmt"
)

type User struct {
	FirstName string
	LastName string
}

func main() {
// creating a map
//one way of making a map
//var myOtherMap map[string]string

//conventional way using make keyword

//A [string]string map
myMap := make(map[string]string) 

myMap["cat"] = "meow"
myMap["dog"] = "bark"

log.Println(myMap["cat"])
log.Println(myMap["dog"])

myMap["dog"]="woof"
log.Println(myMap["dog"])

//A [string]int map
myMap2 := make(map[string]int)
myMap2["first"]=1
myMap2["second"]=2
fmt.Println(" ")
log.Println(myMap2["first"],myMap2["second"])

//A [string]struct map
user := make(map[string]User)

someone := User {
	FirstName: "Dingle",
	LastName: "Dong",
}

user["someone"] = someone

fmt.Println(" ")
log.Println(user["someone"].FirstName)

}
