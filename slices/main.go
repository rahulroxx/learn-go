package main
import "fmt"
import "sort" 

func main() { 
	fmt.Println("Welcome to video on slices")
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)
	fruitList = append(fruitList, "mango", "banana")

	highScores := make([]int, 4)

	highScores[0] = 234 
	highScores[1] = 945
	highScores[2] = 465

	fmt.Println(highScores)
	
	highScores = append(highScores, 555, 666, 321)
	fmt.Println(highScores)
	
	sort.Ints(highScores)
	fmt.Println(highScores)
}
