//https://github.com/go-arun/sps/tree/master/Week_00
#include <stdio.h>
#include<math.h>

void main(){
    int limit;
    int number=1;
    int sum=0;

    printf("Show all Odd numbers upto ? [Enter a limt]:");
    scanf("%d",&limit);
    
    while (number <= limit)
    {

        if (number%2 != 0){
            printf ("%d ,",number);
            sum = sum + number;
        }
        number++;
    }
    printf("\nOdd numers till %d are above and its Sum is : %d",limit,sum);
    


}