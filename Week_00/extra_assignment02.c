#include <stdio.h>
void main(){
int total_rows,i,j,k,s;
int space_maintainer;
printf("Enter number of rows for this special pattern :");
scanf("%d",&total_rows);
space_maintainer = total_rows;

for (i=0;i<total_rows;i++){
        j= i+1 ;// total elements in particular row
        k=1;
        for (s = space_maintainer--;s >0;s--){//maintain starting postion 
            printf(" ");
        }
        while (j>0)
        {
            if (j-1 == 0){//Is last element in this row , then print 1
                k=1;
            }
            printf("%d ",k++);
            j--;
        }
        printf("\n");//new line
        
    }
    
}