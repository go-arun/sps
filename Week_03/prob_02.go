package main

import "fmt"

func main()  {
	var usrInp1 int
	var usrInp2 float32
	fmt.Println("Input a First Number:")
	fmt.Scanln(&usrInp1)
	fmt.Println("Input a Second Number:")
	fmt.Scanln(&usrInp2)
	fmt.Println("Sum is :",float32(usrInp1)+usrInp2)
}