//https://github.com/go-arun/sps/tree/master/Week_00
#include <stdio.h>
#include<math.h>

void main(){
int int_array[10];
int array_size;
int temp,i,j;

printf("Enter array Size[Max Value Limitted to 10]:");
scanf("%d",&array_size);

printf("Enter any %d numbers to populate array_1\n",array_size);
for (i=0;i<array_size;i++){
    scanf("%d",&int_array[i]);
}

for (i=0;i<array_size;i++){
    
    for (j=i+1;j<array_size;j++){
        if ( int_array[j]> int_array[i]){
            temp = int_array[i];
            int_array[i] = int_array[j];
            int_array[j] = temp;
        }
    }
}
print_arry(int_array,array_size);
}


int print_arry(int arr[10],int arr_size){
    //printf("inside func");
    int i;
    printf("[");
     for (i=0;i<arr_size;i++){
         printf("%d ",arr[i]);
     }
    printf("]\n");

}