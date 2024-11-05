package main 

import "fmt" 

func main(){ 
	
	hitesh := User{"rahul", "rahulyup@gmail.com", false, 10}

	fmt.Println(hitesh)
}

type User struct { 
	Name string
	Email string 
	Status bool 
	Age int
	
}
