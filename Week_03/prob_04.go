package main
import "fmt"

func main()  {
	var totalMark float32
	fmt.Println("Enter Total Mark:")
	fmt.Scanln(&totalMark)

	if (totalMark >= 50 ){
		fmt.Println(" Result : PASSED")
	}else {
		fmt.Println(" Result : FAILED")
	}
}