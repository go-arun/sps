//https://github.com/go-arun/sps/tree/master/Week_00
#include <stdio.h>
#include<math.h>

void main(){
    int number;
    int upper_limit=1;
    
    
    printf("Enter a number to see the multipplicaiton table : ");
    scanf("%d",&number);

    while (upper_limit <= 10)
    {
        printf("%d X %d  = %d \n",upper_limit,number,upper_limit*number);
        upper_limit++;
    }
    





}