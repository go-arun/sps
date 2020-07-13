//https://github.com/go-arun/sps/tree/master/Week_00
#include <stdio.h>
#include<math.h>

void main(){
int int_array[10];
int array_size;
int i,count=0;

printf("Enter array Size[Max Value Limitted to 10]:");
scanf("%d",&array_size);

printf("Enter any %d numbers to populate array_1\n",array_size);
for (i=0;i<array_size;i++){
    scanf("%d",&int_array[i]);
}

//finding odd
for (i=0;i<array_size;i++){
    if (int_array[i]%2 ==0 ){
        count++;
    }
}
printf("There are %d Even numbers in this list\n",count);
}