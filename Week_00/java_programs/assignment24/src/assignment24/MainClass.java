package assignment24;

import java.util.Scanner;

public class MainClass {
	
	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		MyClass mc = new MyClass();
		
		System.out.println(" 1 - Circle \n 2 - Square \n 3 - Rectangle \n 4 - Triangle \n 0 -EXIT");
		
	for(;;) {	
		System.out.print("Enter Your Choice :");
		int usrChoice = sc.nextInt();
		if (usrChoice == 1) {
			mc.calulateCircleArea();
		}else if (usrChoice == 2) {
			mc.calulateSquareArea();
		}else if (usrChoice == 3) {
			mc.calulateRectangeArea();
		}else if (usrChoice == 4) {
			mc.calulateTriangleArea();
		}else if (usrChoice == 0) {
			System.out.println("Exiting....");
			break;
		}
		System.out.println("-------------------------------------");
	}
		
		
	}

}
/*
Enter your choice
Circle
Square
Rectangle
Triangle
*/


