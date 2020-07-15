#include<stdio.h>

void main(){
    int stud_count;
    printf("Enter Total Number of Students:");
    scanf("%d",&stud_count);
    int i;
    int num_of_handshakes=0;
    for (i=1;i<stud_count;i++){
        num_of_handshakes = num_of_handshakes + i;
    }
printf("Total Number of handshales will be :%d",num_of_handshakes);
printf("\n------------------------------------");
}