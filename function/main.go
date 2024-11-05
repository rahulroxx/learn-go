package main
import "fmt"


func main () { 
	greeter()
	val := proAdder(1, 3)
	fmt.Println(val)	
}

func adder(valOne int, valTwo int) int { 
	return valOne + valTwo
}
	
func proAdder(values ...int) int { 
	total := 0
	for _, val := range values { 
		total += val
	 }
	return total 
}  
func greeter() { 
	fmt.Println("Go Lang")
}
