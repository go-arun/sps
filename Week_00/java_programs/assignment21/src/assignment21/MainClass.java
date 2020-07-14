package assignment21;

import java.util.Scanner;

public class MainClass {
	public static void main(String Arguments[]) {
		
		Scanner sc = new Scanner(System.in);
	
		int Array1[] = new int [5];
		int Array2[] = new int [4]; // to store result
	
		
		for (int i=0;i<5;i++) { //getting values
			System.out.print("Enter Elements of Array1 Position "+i +":");
			Array1[i] = sc.nextInt();
		}
		
		//calculation
		
		for (int i=0;i<4;i++) { 
			Array2[i] = Array1[i] * Array1[i+1];
		}
				
		//Result
		for (int i=0;i<4;i++) { 
			System.out.print(Array2[i] + " ");
	
		}
		
	}

}
