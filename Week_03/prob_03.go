package main

import "fmt"

func main()  {
	var princiAmount,intrRate,totalYears float32
	fmt.Println("Input Principal amount:")
	fmt.Scanln(&princiAmount)
	fmt.Println("Input Interest rate:")
	fmt.Scanln(&intrRate)
	fmt.Println("Input Number of years:")
	fmt.Scanln(&totalYears)

	fmt.Println("Simple Interest is  :", (princiAmount*intrRate*totalYears)/100)
}

