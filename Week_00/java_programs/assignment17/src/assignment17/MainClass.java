package assignment17;

import java.util.Scanner;

public class MainClass {
	
	public static void main(String arg[]) {
				
		Scanner sc = new Scanner(System.in);
		Operation Op = new Operation();
		
		System.out.println("Enter two Numbers to do Calculations");
		float Fnumber = sc.nextFloat();
		float Snumber = sc.nextFloat();
	for(;;) {
		System.out.print("Press |A->Addition |M->Multipplication|S->Substrattion|D->Division|X->Exit");
		String UsrChoice = sc.next();
		
		if (UsrChoice.equals("A")) {
			Op.Addition(Fnumber, Snumber);
		}else if(UsrChoice.equals("M")) {
			Op.Multipplication(Fnumber, Snumber);
		}else if(UsrChoice.equals("S")) {
			Op.Substraction(Fnumber, Snumber);
		}else if(UsrChoice.equals("D")) {
			Op.Division(Fnumber, Snumber);
		}else if(UsrChoice.equals("X")) {
			System.out.println("Program Exiting ..");
			break;
		}else {
			System.out.println("Wrong Input..Please Try Again..");
		}
		
	}
		

	}

}
