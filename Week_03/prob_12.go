package main
import "fmt"

func main() {

	var ArraySize,i,temp int
	fmt.Printf("Please input Array Size:")
	fmt.Scanln(&ArraySize)
	arr1 := make([]int,ArraySize)
	
	for i = 0; i < ArraySize; i++{
		fmt.Printf("Enter Element for the Position[%d] of Array :",i)
		fmt.Scanln(&arr1[i])
	}
	comparingPosition := 0
	for comparingPosition < ArraySize{
		for i = comparingPosition; i < ArraySize-1; i++{ //sorting
			if arr1[comparingPosition] < arr1[i+1] {
			temp = arr1[comparingPosition]
			arr1[comparingPosition] =  arr1[i+1]
			arr1[i+1] = temp
			}
		}
		comparingPosition++
	}
	fmt.Println("Array Sorted",arr1)
	
}
