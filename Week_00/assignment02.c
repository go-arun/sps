#include <stdio.h>
#include <math.h>
void main(){
        int usr_inp_int;
        float usr_inp_float;

        printf("Enter an Interger(Must be an Integer):");
        scanf("%d",&usr_inp_int);
        printf("Enter a Float Number to get the Sum  :");
        scanf("%f",&usr_inp_float);

        printf("Sum of %g + %d is : %g\n",usr_inp_float,usr_inp_int,(usr_inp_float+usr_inp_int));
}