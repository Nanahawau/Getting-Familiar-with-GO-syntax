package main

import "fmt"

func main ()  {

	emails := make(map[string] string)

	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"


	fmt.Println(emails)
	fmt.Println(len(emails))



	//Declare and Assign

	name := map[string] string{"Bob": "boob@gmail.com", "Sharon": "Sharon@gmail.com"}

	name["mike"] = "mike@gmail.com"  


	fmt.Println(name["mike"])
}