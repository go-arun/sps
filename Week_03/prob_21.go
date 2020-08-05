package main

import "fmt"

func main(){
	var arrSize int
	fmt.Print("Enter Array Size : ")
	fmt.Scanln(&arrSize)
	inpArray 	:= make([]int, arrSize)
	resultArray := make([]int, arrSize-1)
	for  i := 0; i < arrSize; i++ {
		fmt.Printf("Enter Element for the Position [%d] : ",i)
		fmt.Scanln(&inpArray[i])
	}
	for j :=0; j < arrSize-1; j++ {
		resultArray[j] = inpArray[j] * inpArray[j+1]
	}
	fmt.Printf("Result Array given below ")
	fmt.Println(resultArray)
}