package main

import (
	"log"
	// f "first/deleteit"
)

/*
*in go we use slices is more like  array
 */

type Use struct {
	FirstName string
	LastName  string
}

func main() {
	var mySlice []string
	var sliceNumber []int

	// struct type slice

	var userName []Use

	use := Use{
		FirstName: "Asad",
		LastName:  "abaro Asad",
	}

	userName = append(userName, use)

	log.Println(userName)

	mySlice = append(mySlice, "Rahat")
	mySlice = append(mySlice, "Akash")

	sliceNumber = append(sliceNumber, 10)

	log.Println(mySlice, sliceNumber)

}
