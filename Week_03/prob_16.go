package main
import(
	"fmt"
) 
func main() {
	var usrInp,i int
	fmt.Print("Enter Number to Check whether it is a Prime :")
	fmt.Scanln(&usrInp)

	for i = 2 ; i < usrInp; i++{
		if (usrInp%i == 0){
			break
		}
	}
	if (i == usrInp){
		fmt.Println("IS PRIME")
	}else{
		fmt.Println("IS NOT PRIME")
	}
}