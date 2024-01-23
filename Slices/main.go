package main

import (
	"sort"
	"log"
)

func main() {

	//slice of a string type
	var myStringSlice []string
	myStringSlice = append(myStringSlice,"Dingle")
	myStringSlice = append(myStringSlice,"Dong")
	myStringSlice = append(myStringSlice,"jingle")

	log.Println(myStringSlice)
        
	//slice of int type
	var myIntSlice []int
	myIntSlice = append(myIntSlice, 5)
	myIntSlice = append(myIntSlice, 3)
	myIntSlice = append(myIntSlice, 4)
	myIntSlice = append(myIntSlice, 2)
        myIntSlice = append(myIntSlice, 1)

	log.Println(myIntSlice)

	sort.Ints(myIntSlice)

	log.Println(myIntSlice)

	//shorthand decleration of a slice
	numbers := []int{1,2,3,4,5,6,7,8,9}
	log.Println(numbers)
	log.Println(numbers[3:7])
	log.Println(numbers[4])

	strings := []string{"Dingle","berry","thirty four","hakuna matata"}
	log.Println(strings)
	log.Println(strings[0:2])
	log.Println(strings[2])


}
