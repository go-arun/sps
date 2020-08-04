package main
import "fmt"

func main() {
	var limit,sum int
	
	fmt.Println("Enter Upper Limit:")
	fmt.Scanln(&limit)
	fmt.Print("Sum of Odd numbers upto ",limit," is:")
for (limit >0) {
	
	if(limit%2 != 0){
		sum = sum + limit;
	}
	limit--
}
fmt.Println(sum)
}