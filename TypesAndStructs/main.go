package main

import (
	"log"
	"time"
)

type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

func main() {
	user := User {
		FirstName: "Dingle",
		LastName: "Dong",
	}

	log.Println(user.FirstName,user.LastName,"Birth Date",user.BirthDate)
}
