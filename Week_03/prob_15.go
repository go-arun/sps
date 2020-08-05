package main
import(
	"fmt"
) 
func main() {
	var arraySize,i int
	fmt.Println("Enter two Diamention Array Size :")
	fmt.Scanln(&arraySize)
	twoDslice1 := make([][]int,arraySize) //declaraion only 

	for i = range twoDslice1{ 
		twoDslice1[i] = make([]int,arraySize)

	}
	getArray(twoDslice1,arraySize)

	fmt.Println("--Displaying--")
	displayArray(twoDslice1,arraySize)
}
func getArray(twoD [][]int,arrSize int){
	var j int
	for i := range twoD {
		for j = 0; j < arrSize; j++ {
			fmt.Printf("Enter Element for Positon[%d][%d] : ",i,j)
			fmt.Scanln(&twoD[i][j])
		}
	}
}
func displayArray(twoD [][]int,arrSize int){
	var j int
	for i := range twoD {
		for j = 0; j < arrSize; j++ {
			fmt.Print(twoD[i][j]," ")
		}
		fmt.Println("")//new line
	}
}