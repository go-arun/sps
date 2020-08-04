package main
import "fmt"

func main()  {
	var num,i int
	
	fmt.Println("Enter Any Number:")
	fmt.Scanln(&num)
	
	for i=1; i<=10; i++{
		fmt.Printf ( "%d X %d  =  %d \n",	i,num, i*num)
	}
}