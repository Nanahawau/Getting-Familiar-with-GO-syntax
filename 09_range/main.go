package main

import "fmt"

func main ()  {
	
	ids:= []int {33, 76, 54, 23, 11, 2}

	//loop through ids 
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	} 

	//Not using the index

	for _, id := range ids {
		fmt.Printf("ID: %d\n",  id)
	} 


	//Add ids together 

	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println(sum) 


	//Range with Maps  
	emails := map[string] string {"Bob": "Bob@gmail.com", "Sharon": "sharon@gmail.com"}
	 
	for k, v := range emails {
		fmt.Println(k, v)
	}
}