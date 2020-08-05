package main

import "fmt"

func main(){
	i := 1
	j := 1
	rowSkipper := 0
	for i <= 10{
		rowSkipper = 0
		for rowSkipper < j {
			fmt.Print(" ",i)
			i++
			rowSkipper++
		}
		fmt.Println("") //new Line
		j++
	}
}