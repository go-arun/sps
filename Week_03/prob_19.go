package main

import "fmt"

func main(){
	var annualIncome,tax float32
	fmt.Print("Enter Anuual Icome : ")
	fmt.Scanln(&annualIncome)
	if (annualIncome <= 5000000 && annualIncome > 1000000 ){//Above 10 Lakhs to 50 Lakhs
		tax = annualIncome*30/100
	}else if (annualIncome > 500000){
		tax = annualIncome*20/100
	}else if (annualIncome > 250000){
		tax = annualIncome*5/100
	}
	fmt.Println("Tax Amount is :", tax)
}