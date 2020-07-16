//https://github.com/go-arun/sps/tree/master/Week_00
#include <stdio.h>

void main(){
int limit,i=0,temp,j;
int inp_array[20];

printf("Shot howmany numbers ? :[Maz 20]");
scanf("%d",&limit);

for(i=0;i<limit;i++){
    printf("Enter %d more number/s:",limit-i);
    scanf("%d",&inp_array[i]);
}
//Sorting
for(i=0;i<limit;i++){
    for (j=i+1;j<limit;j++){
        if (inp_array[i] < inp_array[j]){
            temp = inp_array[i];
            inp_array[i] = inp_array[j];
            inp_array[j] = temp;
        }
    }
}
//Display Sorted Array
printf("Result is :\n");
for(i=0;i<limit;i++){
    printf("%d\n",inp_array[i]);
}
}