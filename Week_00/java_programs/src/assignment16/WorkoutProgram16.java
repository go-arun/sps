//
package assignment16;

import java.util.Scanner;

public class WorkoutProgram16 {
	
	public static void main ( String arg[]) {
		int DivisibleCount = 0;
		System.out.println("Enter a number to check whether it is a Prime number or not:");
		Scanner sc= new Scanner(System.in);
		int inp_num= sc.nextInt();
		
		for (int i=1;i<=inp_num;i++) {
			if (inp_num%i == 0 ) {
				DivisibleCount++;
			}
		
		}
		
		if (DivisibleCount > 2) {
			System.out.println("Given Number is NOT Prime");
		}else {
			System.out.println("Given Number is Prime");
		}
			
	}

}
