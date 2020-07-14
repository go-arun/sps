package assignment18;

import java.util.Scanner;

public class MainClass {
		
	public static void main (String arg[]) {
		Scanner sc = new Scanner(System.in);
		float WrittenTestMark,LabTestMark,AssignmentMark;
		System.out.println("Enter Marks to calculate Grade");
		
		System.out.print("Written Test Mark:");
		WrittenTestMark = sc.nextFloat();
		
		System.out.print("Lab Test Mark    :");
		LabTestMark = sc.nextFloat();
		
		System.out.print("Assignment Mark  :");
		AssignmentMark = sc.nextFloat();
		
		//written test counts 70%,lab exams 20%, and assignments 10%.
		System.out.println("Grade of this Student is :"+ 
		(((WrittenTestMark*70)/100) 
		+ ((LabTestMark*20)/100)
		+ ((AssignmentMark*10)/100)));
		
	}

}
