//https://github.com/go-arun/sps/tree/master/Week_00
#include <stdio.h>
#include<math.h>

void main(){
    int option;

    printf("Enter any option(from 1 to 7 ):");
    scanf("%d",&option);

    switch (option){
        case 1:
            printf("Sunday");
            break;
        case 2:
            printf("Monday");
            break;
        case 3:
            printf("Tuesday");
            break;
        case 4:
            printf("Wednesday");
            break;        
        case 5:
            printf("Thursday");
            break;
        case 6:
            printf("Friday");
            break;        
        case 7:
            printf("Saturday");
            break;        
        default:
            printf("invalid Entry Please try Again\n");
            break;
    }


}