package main

import "fmt"

func main ()  {
	a := 5
	//assigning b to a pointer of a 
	//You get a memory address of a 
	b := &a


	fmt.Println(a, b);
	//(* int)
	fmt.Printf("%T\n", b);

	//Use to read value from memory address 
	fmt.Println(*b);

	//Same as 
	fmt.Println(*&a);

	//Change value of a with pointer

	*b = 10;

	fmt.Println(a);


}