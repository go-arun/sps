package main
import "fmt"

func main() {

	var ArraySize,i,count int
	fmt.Printf("Please input Array Size:")
	fmt.Scanln(&ArraySize)

	arr1 := make([]int,ArraySize)

	for i = 0; i < ArraySize; i++{
		fmt.Printf("Enter Element for the Position[%d] of Array :",i)
		fmt.Scanln(&arr1[i])
	}
	for i = 0; i < ArraySize; i++{ //counting odd nums
		if (arr1[i]%2 == 0 ){
			count++
		}
	}
	fmt.Println(arr1)
	fmt.Println("Total Even numbers in above  array is  :",count)
}
