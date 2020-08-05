package main
import(
	"fmt"
) 

type inputs struct{
	firstNum,secndNum int
}

func main() {
	var usrInp inputs
	var operationSelector string
	fmt.Println("Enter two Numbers to do all arithmetic Calculations")
	fmt.Print("Enter First Number  : ")
	fmt.Scanln(&usrInp.firstNum)
	fmt.Print("Enter Second Number : ")
	fmt.Scanln(&usrInp.secndNum)
for (operationSelector != "q"){
	fmt.Println("Choose the Operaton you want to Perform [+] [-] [*] [/] q -> Exit")
	fmt.Scanln(&operationSelector)
	switch operationSelector {
	case "+":
		fmt.Println(usrInp.addition())
	case "-":
		fmt.Println(usrInp.substraction())
	case "*":
		fmt.Println(usrInp.mutliplication())
	case "/":
		fmt.Println(usrInp.division())
	case "q":
		fmt.Println("Thank You! Exiting ..")		
	default:
		fmt.Println("Invalid Entry Plase try again..")
	}
}
	

}
func (inp inputs) addition() ( result int){
	 result = inp.firstNum + inp.secndNum
	 return
}
func (inp inputs) mutliplication() ( result int){
	result = inp.firstNum * inp.secndNum
	return
}
func (inp inputs) division() ( result int){
	result = inp.firstNum /inp.secndNum
	return
}
func (inp inputs) substraction() ( result int){
	result = inp.firstNum - inp.secndNum
	return
}

