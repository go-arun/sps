package assignment22;

import java.util.Scanner;

public class MainClass {
	
	public static void main (String Arg[]) {
		
		Scanner scn = new Scanner(System.in);
		System.out.print("Enter Size of array : ");
		int arraySize = scn.nextInt();
	
		
		int Array1[][] = new int[10][10];
		int Array2[][] = new int[10][10];

		System.out.println("Fill Elements of First Array");
		getArray(Array1,arraySize);
		System.out.println("Fill Elements of Second Array");
		getArray(Array2,arraySize);
		
		AddArray(Array1, Array2,arraySize);
		System.out.println("Two Arrays added , Result as follows:");
		displayArray(Array1,arraySize); // Result is stored in Array1,so displaying it.

	}
	
	public static void getArray(int Arr[][],int arrSize) {
		Scanner sc = new Scanner(System.in);
		for (int i=0;i<arrSize;i++) {
			for (int j=0;j<arrSize;j++) {
				System.out.print("Enter Element for position["+ i +"," + j +"] :");
				Arr[i][j] = sc.nextInt();
			}
			
		}
		
		
	}
	public static void displayArray(int Arr[][],int arrSize) {
		for (int i=0;i<arrSize;i++) {
			for (int j=0;j<arrSize;j++) {
				System.out.print(Arr[i][j] + " ");
				
			}
			System.out.println();//new line
	
		}
		
	}
	public static void AddArray(int Arr1[][],int Arr2[][],int arrSize) {
		for (int i=0;i<arrSize;i++) {
			for (int j=0;j<arrSize;j++) {
				Arr1[i][j] = Arr1[i][j] + Arr2[i][j];//result storing in first array itself
				
			}
			System.out.println();//new line
	
		}
		
	}
}
