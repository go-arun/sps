//https://github.com/go-arun/sps/tree/master/Week_00
#include <stdio.h>
#include<math.h>
#include <string.h>

void main(){
int i,j;
int array_1[3][3];
int array_2[3][3];

for (i=0;i<3;i++){ //fill array1
    for (j=0;j<3;j++){
        printf("Enter value for postion[%d][%d] of Array1 :",i,j);
        scanf("%d",&array_1[i][j]);
    }
}

for (i=0;i<3;i++){ //fill array2
    for (j=0;j<3;j++){
        printf("Enter value for postion[%d][%d] of Array2 :",i,j);
        scanf("%d",&array_2[i][j]);
    }
}
printf("Array1:\n");
print_array(array_1);
printf("Array2:\n");
print_array(array_2);

//Sum storring in array1 itself
for (i=0;i<3;i++){
    for (j=0;j<3;j++){
      array_1[i][j] = array_1[i][j] + array_2[i][j];  
        
    }
}
printf("Sum of these two :\n");

print_array(array_1);

}

int print_array(int arr[3][3]){
    int i,j;
    for (i=0;i<3;i++){ 
        for (j=0;j<3;j++){
            printf ("%d ",arr[i][j]);
        }
        printf("\n");//new line
    }
}
