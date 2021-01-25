package main

import ("fmt"
			"strconv")

//Define person struct. Structs are classes is in Go


type Person struct {
	firstName,middleName,gender, city string
	age int
}

	// Methods in GO:
	//Value receivers and Pointer receivers 

	//Value Receivers 

//Method - Value Receivers: Greeting 
func (person Person) greet() string  {
	return "Hello" + person.firstName + ", my age is" + strconv.Itoa(person.age)
}


//Pointer Receiver
 func (person *Person) hasBirthday() {
	 person.age++
 }



func main ()  {
	fmt.Println("Hello World")

	personAli := Person {firstName: "Ali",
							middleName: "Mohammed", age: 30, 
						gender: "F", 
						 city: "Lagos"}

						 fmt.Println(personAli);


	personJ := Person {"Jums", "Adeks","F", "Tanke", 20}

	fmt.Println(personJ);

	fmt.Println(personJ.firstName);
	







						  
}