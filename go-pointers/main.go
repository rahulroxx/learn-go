package main

import "fmt"

func main() { 
	fmt.Println("Welcome to a class on pointers")
	
	myNumber := 23

	var ptr = &myNumber
	fmt.Println(ptr)
}
