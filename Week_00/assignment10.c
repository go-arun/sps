//https://github.com/go-arun/sps/tree/master/Week_00
#include <stdio.h>
#include<math.h>

void main(){
int array_1[10];
int array_2[10];
int array_size;
int i,temp;

printf("Enter array Size[Max Value Limitted to 10]:");
scanf("%d",&array_size);

printf("Enter any %d numbers to populate array_1\n",array_size);
for (i=0;i<array_size;i++){
    scanf("%d",&array_1[i]);
}
printf("Enter any %d numbers to populate array_2\n",array_size);
for (i=0;i<array_size;i++){
    scanf("%d",&array_2[i]);
}
printf("User defined arrays shown below\nArray1:");
print_arry(array_1,array_size);
printf("\nArray2:");
print_arry(array_2,array_size);

//interchaning 
for(i=0;i<array_size;i++){
    temp = array_1[i];
    array_1[i] = array_2[i];
    array_2[i] = temp;
}

printf("Interchanged array elements[Result shown below]\nArray1:");
print_arry(array_1,array_size);
printf("Array2:");
print_arry(array_2,array_size);

}

//function to print aray
int print_arry(int arr[10],int arr_size){
    //printf("inside func");
    int i;
    printf("[");
     for (i=0;i<arr_size;i++){
         printf("%d ",arr[i]);
     }
    printf("]\n");

}
