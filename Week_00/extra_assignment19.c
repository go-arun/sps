#include<stdio.h>

void main(){
    int parcel_weight,chargable_weight;
    int chargable_amt=200; // minimum amount
    printf("Enter Parcel Weight(in Grams :");
    scanf("%d",&parcel_weight);
   
    if (parcel_weight > 500){ 
        chargable_weight = parcel_weight - 500;
        if (chargable_weight >= 500){ // Upto next 500 is free, so this condition
            chargable_amt = chargable_amt + ((chargable_weight/500) * 170);
        }
    }
printf (" Postage charge for this Parcel is :%d \n",chargable_amt);
 }
    /*
    less than 500gm or eq
    
    ual to 500gm then the parcel charge will be Rs. 200,
     Otherwise there is an additional charge of Rs.170 per each extra 500gm
     */
   