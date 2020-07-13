//https://github.com/go-arun/sps/tree/master/Week_00
#include <stdio.h>
#include<math.h>
#include <string.h>

void main(){
int i,j;
int array[3][3];

get_array(array);
display_array(array);

}

//functions 
int display_array(int arr[3][3]){
    int i,j;
    for (i=0;i<3;i++){ 
        for (j=0;j<3;j++){
            printf ("%d ",arr[i][j]);
        }
        printf("\n");//new line
    }
}

int get_array(int arr[3][3]){
    int i,j;
    for (i=0;i<3;i++){ //fill array2
    for (j=0;j<3;j++){
        printf("Enter value for postion[%d][%d] of Array :",i,j);
        scanf("%d",&arr[i][j]);
    }
}
}