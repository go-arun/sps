#include <stdio.h>
#include <math.h>

void main (){
    int princ_amt;
    float no_years;
    float inter_rate;
    float simp_interest;

    printf("Enter Principal amount       :");
    scanf("%d",&princ_amt);
    printf("Enter Total number of Years  :");
    scanf("%f",&no_years);
    printf("Enter Interest Rate          :");
    scanf("%f",&inter_rate);
    
    printf("Simple interest              :%g \n",(princ_amt*inter_rate*no_years)/100);
    //Formula: SI=(P*R*n)/100)

}