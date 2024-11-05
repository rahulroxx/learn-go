package main
import "math/rand"
import "fmt"
import "time"

func main() { 
	fmt.Println("Switch and case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNumber)
	
	switch diceNumber { 
		case 1:
			fmt.Println("Dice value is 1 and you can open") 
		case 2:
			fmt.Println("Dice value is 2 and you can run two spots")

		default:	
			fmt.Println("what was that")


	}
} 
