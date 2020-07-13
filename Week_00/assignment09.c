//https://github.com/go-arun/sps/tree/master/Week_00
#include <stdio.h>
#include<math.h>

void main(){
    
    int i=1;
    int j=1;

    for (i=1;i <= 10;i++){
        for (j=1; j<= i;j++){
            printf("%d ",j);
            
        }
       printf("\n");//new line 
        
    }
}