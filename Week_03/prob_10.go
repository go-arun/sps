package main
import "fmt"

func main() {

	var ArraySize,i,temp int
	fmt.Printf("Please input Array Size:")
	fmt.Scanln(&ArraySize)

	arr1 := make([]int,ArraySize)
	arr2 := make([]int,ArraySize)
	for i = 0; i < ArraySize; i++{
		fmt.Printf("Enter Element for the Position[%d] of Array1 :",i)
		fmt.Scanln(&arr1[i])
	}
	fmt.Println("------------------------------------------------")
	for i = 0; i < ArraySize; i++{
		fmt.Printf("Enter Element for the Position[%d] of Array2 :",i)
		fmt.Scanln(&arr2[i])
	}
	fmt.Println("Before interChanging")
	fmt.Println(arr1,arr2)
	//swaping
	for i=0; i<ArraySize ;i++{
		temp = arr1[i]
		arr1[i] = arr2[i]
		arr2[i] = temp
	}
	fmt.Println("After interChanging")
	fmt.Println(arr1,arr2)
}