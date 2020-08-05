package main

import "fmt"

func main (){
	var arraySize int 
	fmt.Print("Enter size of Array : ")
	fmt.Scanln(&arraySize)

	inpArray1 := make([][]int,arraySize )
	inpArray2 := make([][]int,arraySize )
	for i := range inpArray1 {//making 2 diamention
		inpArray1[i] = make([]int, arraySize)
		inpArray2[i] = make([]int,arraySize )
	}

	fmt.Println("Getting Values for First   Array")
	getArray(inpArray1,arraySize)
	fmt.Println("Getting Values for Second  Array")
	getArray(inpArray2,arraySize)
	addArray(inpArray1,inpArray2,arraySize)
	fmt.Println("Printing Result Array")
	displayArray(inpArray1,arraySize)

}
func addArray(array1 [][]int,array2 [][]int,arrSize int){
	for i := range array1{	
		for j := 0; j < arrSize; j++ {
			array1[i][j] = array1[i][j] + array2[i][j]
		}
	
	}
}
func displayArray(array [][]int, arrSize int ){
	for i := range array{	
		for j := 0; j < arrSize; j++ {
			fmt.Print(" ",array[i][j])
		}
		fmt.Println("")//new line
	}
}

func getArray(array [][]int, arrSize int ){
	for i := range array{	
		for j := 0; j < arrSize; j++ {
			fmt.Printf("Enter Element for the Positon[%d][%d] : ",i,j)
			fmt.Scanln(&array[i][j])
		}
		
	}
}
