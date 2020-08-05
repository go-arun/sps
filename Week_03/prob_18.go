package main

import "fmt"

type marksScored struct{
	wittenTest,labExam,assignMent float32
}
func main(){
	var student marksScored
	fmt.Println("Provide the Marsk scored in Each Exams")
	fmt.Print(" Written Test :")
	fmt.Scanln(&student.wittenTest)
	fmt.Print(" Lab Exams    :")
	fmt.Scanln(&student.labExam)
	fmt.Print(" Assignment   :")
	fmt.Scanln(&student.assignMent)
	fmt.Print("Grade of this student is : ")
	fmt.Println(student.wittenTest*70/100 + student.labExam*20/100 + student.assignMent*10/100)
}

