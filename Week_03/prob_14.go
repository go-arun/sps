package main
import(
	"fmt"
) 
func main() {
	var arraySize,j,i int

	fmt.Println("Enter two Diamention Array Size :")
	fmt.Scanln(&arraySize)
	twoDslice1 := make([][]int,arraySize) //declaraion only 
	twoDslice2 := make([][]int,arraySize)

	for i = range twoDslice1{ 
		twoDslice1[i] = make([]int,arraySize)
		twoDslice2[i] = make([]int,arraySize)
	}

	for i = range twoDslice1{ 
		for j = 0; j<arraySize; j++ {
			fmt.Printf("Enter Element for Array1 Positon[%d][%d] : ",i,j)
			fmt.Scanln(&twoDslice1[i][j])
		}
	}
	fmt.Println("---------------------------------")
	for i = range twoDslice2{ 
		for j = 0; j<arraySize; j++ {
			fmt.Printf("Enter Element for Array2 Positon[%d][%d] : ",i,j)
			fmt.Scanln(&twoDslice2[i][j])
		}
	}
	for i = range twoDslice2{ 
		for j = 0; j<arraySize; j++ {
			twoDslice1[i][j] = twoDslice1[i][j] + twoDslice2[i][j]
		}
	}
	fmt.Println("--------------Result-------------")

	for i = range twoDslice2{ 
		for j = 0; j<arraySize; j++ {
			fmt.Print(twoDslice1[i][j] ," ")
		}
		fmt.Println("")//new line
	}
}