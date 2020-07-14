package assignment24;

import java.util.Scanner;

public class Area {
	Scanner sc = new Scanner(System.in);
	
	void calulateCircleArea() { //A = π r²)
		System.out.print("Enter Radiou of Circle :");
		float circleRadious = sc.nextFloat();
		System.out.println("Area of Circle is :" + (3.141592653589793 * circleRadious * circleRadious ));
		
	}
	void calulateSquareArea() { //A=a2
		System.out.print("Enter One Side Length of Square :");
		float squareLength = sc.nextFloat();
		System.out.println("Area of Square is :" + ( squareLength * squareLength ));
		
	}
	void calulateRectangeArea() { //A=w*l
		System.out.print("Enter Lenth of Rectangle :");
		float RectangelLength = sc.nextFloat();
		System.out.print("Enter Width of Rectangle :");
		float RectangeWidth = sc.nextFloat();
		System.out.println("Area of Rectangle is :" + ( RectangelLength * RectangeWidth ));
		
	}
	void calulateTriangleArea() { // base * height/ 2
		System.out.print("Enter base Length of Triangle :");
		float baseLength = sc.nextFloat();
		System.out.print("Enter One height of Triangle  :");
		float heightLength = sc.nextFloat();
		System.out.println("Area of Square is :" + (( baseLength * heightLength )/2));
		
	}
	
}



/*circle()
square()
rectangle()
triangle()
*/