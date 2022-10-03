package main

import (
	"log"
	"time"
)

//learning the type and struct in go
/*
if we have to store to many variables for any perticuler purpose then we can use type or struct
*/

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BrithDate   time.Time
}

func main() {
	user := User{

		FirstName: "rahat",
		LastName:  "Akash",
	}

	log.Println(user.FirstName, user.LastName)
}
