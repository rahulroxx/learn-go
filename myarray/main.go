package main

import "fmt"

func main()  { 
	fmt.Println("Welcome to array in garbage")
	
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[2] = "Apple"
	fruitList[3] = "Ginger"
//	fruitList[4] = "What"

	fmt.Println("fruitlist length", len(fruitList))

	fmt.Println(fruitList)
	

}
