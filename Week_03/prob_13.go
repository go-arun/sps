package main
import(
	"fmt"
	//"strings"

) 
func main() {
	var word string
	fmt.Printf("Please Enter Word to Check:")
	fmt.Scanln(&word)
	halfLength := (len(word))/2
	lastArrayPosition := len(word)-1
	machedPositions := 0

	for i :=0 ; i< halfLength; i++{
		if (word[i] == word[lastArrayPosition]){//comparing ascii value
			machedPositions++
		}
		lastArrayPosition--
	}
	if (machedPositions == halfLength){
		fmt.Println("Is Palindrome ")
	}else{
		fmt.Println("Is NOT a Palindrome ")

	}
}