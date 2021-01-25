package main

import "fmt"

func main ()  {
	//Declare
	var fruitArr [2]string

	//Assign
	fruitArr[0] = "Apple"
	fruitArr[1] = "Pines"


	//Declare Assign in one Line

	fruitAr := [2]string{"Apple", "Orange"}


	fruitSlice := []string {"Apple", "Orange", "Grape", "Cherry"}


	fmt.Println (len(fruitAr), len(fruitArr), len(fruitSlice))


}