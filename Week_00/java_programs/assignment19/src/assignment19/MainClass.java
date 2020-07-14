package assignment19;

import java.util.Scanner;

public class MainClass {

	public static void main(String Argument[]) {
		Scanner sc = new Scanner(System.in);

		System.out.print("Enter Annual Income to calculate Tax :");
		float AnnualIncome = sc.nextFloat();
		
		if (AnnualIncome < 250000) {
			System.out.println("TAX Payable : Excempted!!");
		}else if (AnnualIncome <= 500000) {
			
			System.out.println("TAX Payable [5%]: "+ ((AnnualIncome*5)/100));
		}else if (AnnualIncome <= 1000000) {
			
			System.out.println("TAX Payable [20%]: "+ ((AnnualIncome*20)/100));
		}else if (AnnualIncome <= 5000000) {
			
			System.out.println("TAX Payable [30%] : "+ ((AnnualIncome*30)/100));
		}else {
			System.out.println(" Please verify input!!");
		}
		
	}
}

