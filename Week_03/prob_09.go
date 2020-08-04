package main
import "fmt"

func main() {
	var limit,j int
	
	fmt.Println("Enter Req.Lines:")
	fmt.Scanln(&limit)

for i :=1 ;i<=limit ;i++ {
	for j= 1; j <= i; j++ {

		fmt.Printf(" %d",j)
	}
	fmt.Printf("\n")//new line
}

}