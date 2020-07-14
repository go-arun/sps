package assignment23;

import java.util.Scanner;

public class MainClass {
	
	public static void main(String a[]) {
		System.out.print("Enter Array Size :");
		Scanner sc = new Scanner(System.in);
		
		int sizeOfArray = sc.nextInt();
		ManageArray mr = new ManageArray();
		mr.getArray(sizeOfArray);
		mr.showArray(sizeOfArray);
		

	}

}
