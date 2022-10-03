package main

import "log"

/*
*myStruct has a function associated with it
 */
type myStruct struct {
	FirstName string
}

/*
*here we have diclear a function which can access my type struct
* and return a string value
 */
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {

	var myVar myStruct
	myVar.FirstName = "rahat"

	myVar2 := myStruct{
		FirstName: "Akash",
	}

	log.Println("myVar is set to ", myVar.printFirstName())
	log.Println("myVar is set to ", myVar2.printFirstName())

}
