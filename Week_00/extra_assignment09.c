//https://github.com/go-arun/sps/tree/master/Week_00
#include <stdio.h>

void main(){
char usr_inp[1];
int ascii_val;

printf("Enter characters one by one to know the ASCII equiviant[X ->Exit]\n");

while (1)
{
    scanf("%c",&usr_inp[0]);
    ascii_val = ("%d",usr_inp[0]);
    if (ascii_val == 10 ){continue;}// to avoid showing ascii val of Enter Key
    printf("ASCII Val of %c is %d\n",usr_inp[0],ascii_val);
    if (usr_inp[0] == 'X' || usr_inp[0] == 'x'){
        printf("Exiting..\n");
        break;
    }

}

}