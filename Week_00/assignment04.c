#include <stdio.h>
#include <math.h>

void main(){
    float stud_mark;

    printf("Enter Mark :");
    scanf("%f",&stud_mark);

    if (stud_mark >= 50){
        printf("Result : PASSED\n");
    }else{
        printf("Result     : FAILED\n");
    }

}