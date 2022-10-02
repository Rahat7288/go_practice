package main

import "log"

func main(){
	var name string
	name = "Rahat"
	log.Println("may name is ", name)

	changeName(&name)
	log.Println("after the function call name is", name)
}

func changeName(s *string){
	newName := "akash"
	*s = newName


}