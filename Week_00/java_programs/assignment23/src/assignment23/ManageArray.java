package assignment23;

import java.util.Scanner;

public class ManageArray {

	int [][]Array;
		
	void getArray(int arSize) { //Method 1
		Array = new int[arSize][arSize];
		//int Array[][] = new int[arSize][arSize];
		Scanner sc = new Scanner(System.in);
		for (int i=0;i<arSize;i++) {
			for (int j=0;j<arSize;j++) {
				System.out.print("Enter Element for the position ["+ i + "," + j + "]:");
				Array[i][j] = sc.nextInt();
			}
		}
	}

	void showArray(int arSize) { //Method 2
		Scanner sc = new Scanner(System.in);
		System.out.println("Array elements are:");
		for (int i=0;i<arSize;i++) {
			for (int j=0;j<arSize;j++) {
				System.out.print(Array[i][j] + " ");
			}
			System.out.println();//for new line
		}
	}


}
