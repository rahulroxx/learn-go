package main
import "fmt"

func main () { 
	days := []string{"Sunday", "Tuesday", "Wednesday"}
	for _, day := range days { 
		fmt.Println(day)
	}

	rogueValue := 1

	for rogueValue < 10 { 

		fmt.Println("Value is: ", rogueValue) 
		rogueValue++  
 	}
}
