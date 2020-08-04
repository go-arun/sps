package main
import "fmt"

func main()  {
	var totalMark float32
	
	fmt.Println("Enter Total Mark:")
	fmt.Scanln(&totalMark)

	if (totalMark < 50 ) {
		fmt.Println("RESULT: FAILED")
	}else if(totalMark < 60 ){
		fmt.Println("GRADE E")
	}else if(totalMark < 70 ){
		fmt.Println("GRADE D")
	}else if(totalMark < 80 ){
		fmt.Println("GRADE C")
	}else if(totalMark < 90 ){
		fmt.Println("GRADE B")
	}else{
		fmt.Println("GRADE A")
	}
}