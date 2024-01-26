package main

import (
	"log"
)

func main() {

	var isTrue bool
	isTrue = true


	if isTrue{
		log.Println("isTrue is", isTrue)
	}else {
		log.Println("isTrue is", isTrue)
	}

        //strings and if-else 
	cat := "cat"
	if cat == "cat" {
		log.Println("Cat is a Cat")
	} else {
	        log.Println("Cat is not a cat")	
	}

	//numbers and if-else if
	myNum := 100

	if myNum > 99 && isTrue  {
		log.Println("myNum is greater than 99 and isTrue is set to True")
	} else if myNum <100 && isTrue {
		log.Println("myNum is less than 100 and isTrue is set to True")
	}else if myNum == 101 || isTrue{
		log.Println("myNum is exactly 101 and is set to True")
	}else if myNum > 1000 && isTrue==false{
		log.Println("myNum is greater than 1000 and isTrue is set to False")
	}else {
		log.Println("None of the cases")
	}

	//switch statements
	myCat := "fish"

	switch myCat {
		case "cat" :
			log.Println("Cat is set to cat")
		case "dog" :
			log.Println("Cat is set to dog")
		case "fish":
			log.Println("cat is set to fish")
		default :
		        log.Println("cat is set to something else")
		}
}
