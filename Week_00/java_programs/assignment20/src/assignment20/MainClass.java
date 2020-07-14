package assignment20;

public class MainClass {
	
public static void main (String Args[]) {
		
	int row =1;
	int RetainRowNumber = 1;
	for (int number=1;number<=10;) {
		for (;row>0;row--) {
			System.out.print( (number++) + " ");
		}
		System.out.println();//new line
		RetainRowNumber++;
		row = RetainRowNumber;
	}
	
}
}
