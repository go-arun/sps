#include <stdio.h>

void main(){
    int a;
    int b;
    int *a_addr;
    int *b_addr;
    int tmp;
    a = 10;
    b = 20;
    printf("[Before Swapping]value of a = %d and b = %d\n",a,b);
    a_addr = &a;
    b_addr = &b;

    tmp = a;
    *a_addr = b;
    *b_addr = tmp;
    printf("[After Swapping ]value of a = %d and b = %d\n",a,b);
}