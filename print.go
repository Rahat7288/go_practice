package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	/*
	*map[key inside this field] value of this key/return type
	 */

	myMap := make(map[string]string)

	myMap["cat"] = "Pande??"

	log.Println(myMap["cat"])

	//map using the struc type

	myMap2 := make(map[string]User)

	me := User{
		FirstName: "Rahat",
		LastName:  "Akash",
	}
	myMap2["me"] = me

	log.Println(myMap2["me"].FirstName)

}
