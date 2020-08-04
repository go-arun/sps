package main

import "fmt"

func main()  {
	var inpString string
	fmt.Println("Input a String!")

	fmt.Scanln(&inpString)
	fmt.Println("You Entered :" + inpString)
}