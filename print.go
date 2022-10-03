package main

import "log"

/*
*in go we use slices is more like  array
 */

 type Use struct{
	FirstName string
	LastName  string
 }
func main() {
	var mySlice []string
	var sliceNumber []int

	

	mySlice = append(mySlice, "Rahat")
	mySlice = append(mySlice, "Akash")

	sliceNumber = append(sliceNumber, 10)

	log.Println(mySlice, sliceNumber)

}
