//https://github.com/go-arun/sps/tree/master/Week_00
#include <stdio.h>
#include <math.h>

void main(){

    float mark;
    printf("Enter Mark :");
    scanf("%f",&mark);

    if (mark >= 90){
        printf("Grade is : A\n");
    }else if (mark >=80 ){
        printf("Grade is : B\n");
    }else if (mark >= 70){
        printf("Grade is : C\n");
        
    }else if (mark >= 60){
        printf("Grade is : D\n");
    }else if (mark >= 50){
        printf("Grade is : E\n");
    }else{
        printf("Result   : FAILED \n");
    }
    

}